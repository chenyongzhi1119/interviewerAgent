package memory

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	_ "modernc.org/sqlite"
)

// Storage 存储接口，可换 MySQL/PostgreSQL。
type Storage interface {
	GetUserProfile(ctx context.Context, userID string) (*UserProfile, error)
	UpsertUserProfile(ctx context.Context, p *UserProfile) error
	AddQuestionRecord(ctx context.Context, r *QuestionRecord) error
	GetWeaknesses(ctx context.Context, userID string) ([]*WeaknessRecord, error)
	UpsertWeakness(ctx context.Context, w *WeaknessRecord) error
	RemoveWeakness(ctx context.Context, userID, tag string) error
	GetRecentRecords(ctx context.Context, userID string, limit int) ([]*QuestionRecord, error)
	PruneExpiredWeaknesses(ctx context.Context) error
}

// SQLiteStorage 基于 SQLite 的本地持久化（生产可换 MySQL，接口不变）。
type SQLiteStorage struct {
	db *sql.DB
}

func NewSQLiteStorage(path string) (*SQLiteStorage, error) {
	db, err := sql.Open("sqlite", path)
	if err != nil {
		return nil, fmt.Errorf("open sqlite: %w", err)
	}
	db.SetMaxOpenConns(1) // SQLite 单写
	s := &SQLiteStorage{db: db}
	return s, s.migrate()
}

func (s *SQLiteStorage) migrate() error {
	_, err := s.db.Exec(`
	CREATE TABLE IF NOT EXISTS user_profiles (
		user_id       TEXT PRIMARY KEY,
		interview_count INT DEFAULT 0,
		overall_level   INT DEFAULT 1,
		total_questions INT DEFAULT 0,
		avg_score       REAL DEFAULT 0,
		created_at      DATETIME,
		updated_at      DATETIME
	);

	CREATE TABLE IF NOT EXISTS weakness_records (
		user_id          TEXT,
		tag              TEXT,
		weakness_score   REAL DEFAULT 0,
		occurrence_count INT  DEFAULT 0,
		last_seen        DATETIME,
		expires_at       DATETIME,
		PRIMARY KEY (user_id, tag)
	);

	CREATE TABLE IF NOT EXISTS question_records (
		id          INTEGER PRIMARY KEY AUTOINCREMENT,
		session_id  TEXT,
		user_id     TEXT,
		phase       TEXT,
		difficulty  INT,
		tags        TEXT,
		question    TEXT,
		answer      TEXT,
		score       REAL DEFAULT 0,
		created_at  DATETIME
	);
	CREATE INDEX IF NOT EXISTS idx_qr_user ON question_records(user_id, created_at);
	`)
	return err
}

func (s *SQLiteStorage) GetUserProfile(ctx context.Context, userID string) (*UserProfile, error) {
	row := s.db.QueryRowContext(ctx,
		`SELECT user_id,interview_count,overall_level,total_questions,avg_score,created_at,updated_at
		 FROM user_profiles WHERE user_id=?`, userID)
	p := &UserProfile{}
	err := row.Scan(&p.UserID, &p.InterviewCount, &p.OverallLevel,
		&p.TotalQuestions, &p.AvgScore, &p.CreatedAt, &p.UpdatedAt)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	return p, err
}

func (s *SQLiteStorage) UpsertUserProfile(ctx context.Context, p *UserProfile) error {
	_, err := s.db.ExecContext(ctx, `
	INSERT INTO user_profiles(user_id,interview_count,overall_level,total_questions,avg_score,created_at,updated_at)
	VALUES(?,?,?,?,?,?,?)
	ON CONFLICT(user_id) DO UPDATE SET
		interview_count=excluded.interview_count,
		overall_level=excluded.overall_level,
		total_questions=excluded.total_questions,
		avg_score=excluded.avg_score,
		updated_at=excluded.updated_at`,
		p.UserID, p.InterviewCount, p.OverallLevel,
		p.TotalQuestions, p.AvgScore, p.CreatedAt, p.UpdatedAt)
	return err
}

func (s *SQLiteStorage) AddQuestionRecord(ctx context.Context, r *QuestionRecord) error {
	tags, _ := json.Marshal(r.Tags)
	_, err := s.db.ExecContext(ctx, `
	INSERT INTO question_records(session_id,user_id,phase,difficulty,tags,question,answer,score,created_at)
	VALUES(?,?,?,?,?,?,?,?,?)`,
		r.SessionID, r.UserID, r.Phase, r.Difficulty,
		string(tags), r.Question, r.Answer, r.Score, r.CreatedAt)
	return err
}

func (s *SQLiteStorage) GetWeaknesses(ctx context.Context, userID string) ([]*WeaknessRecord, error) {
	rows, err := s.db.QueryContext(ctx,
		`SELECT user_id,tag,weakness_score,occurrence_count,last_seen,expires_at
		 FROM weakness_records WHERE user_id=? AND expires_at > ? ORDER BY weakness_score DESC`,
		userID, time.Now())
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var list []*WeaknessRecord
	for rows.Next() {
		w := &WeaknessRecord{}
		if err := rows.Scan(&w.UserID, &w.Tag, &w.WeaknessScore,
			&w.OccurrenceCount, &w.LastSeen, &w.ExpiresAt); err != nil {
			return nil, err
		}
		list = append(list, w)
	}
	return list, rows.Err()
}

func (s *SQLiteStorage) UpsertWeakness(ctx context.Context, w *WeaknessRecord) error {
	_, err := s.db.ExecContext(ctx, `
	INSERT INTO weakness_records(user_id,tag,weakness_score,occurrence_count,last_seen,expires_at)
	VALUES(?,?,?,?,?,?)
	ON CONFLICT(user_id,tag) DO UPDATE SET
		weakness_score=excluded.weakness_score,
		occurrence_count=excluded.occurrence_count,
		last_seen=excluded.last_seen,
		expires_at=excluded.expires_at`,
		w.UserID, w.Tag, w.WeaknessScore,
		w.OccurrenceCount, w.LastSeen, w.ExpiresAt)
	return err
}

func (s *SQLiteStorage) RemoveWeakness(ctx context.Context, userID, tag string) error {
	_, err := s.db.ExecContext(ctx, `DELETE FROM weakness_records WHERE user_id=? AND tag=?`, userID, tag)
	return err
}

func (s *SQLiteStorage) GetRecentRecords(ctx context.Context, userID string, limit int) ([]*QuestionRecord, error) {
	rows, err := s.db.QueryContext(ctx,
		`SELECT id,session_id,user_id,phase,difficulty,tags,question,answer,score,created_at
		 FROM question_records WHERE user_id=? ORDER BY created_at DESC LIMIT ?`,
		userID, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var list []*QuestionRecord
	for rows.Next() {
		r := &QuestionRecord{}
		var tagsJSON string
		if err := rows.Scan(&r.ID, &r.SessionID, &r.UserID, &r.Phase, &r.Difficulty,
			&tagsJSON, &r.Question, &r.Answer, &r.Score, &r.CreatedAt); err != nil {
			return nil, err
		}
		json.Unmarshal([]byte(tagsJSON), &r.Tags)
		list = append(list, r)
	}
	return list, rows.Err()
}

func (s *SQLiteStorage) PruneExpiredWeaknesses(ctx context.Context) error {
	_, err := s.db.ExecContext(ctx, `DELETE FROM weakness_records WHERE expires_at <= ?`, time.Now())
	return err
}

// TopWeakTags 返回最薄弱的 N 个标签。
func (s *SQLiteStorage) TopWeakTags(ctx context.Context, userID string, n int) ([]string, error) {
	rows, err := s.db.QueryContext(ctx,
		`SELECT tag FROM weakness_records WHERE user_id=? AND expires_at > ?
		 ORDER BY weakness_score DESC LIMIT ?`, userID, time.Now(), n)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var tags []string
	for rows.Next() {
		var t string
		rows.Scan(&t)
		tags = append(tags, t)
	}
	return tags, nil
}

// TagScoreMap 返回用户各标签的历史平均得分。
func (s *SQLiteStorage) TagScoreMap(ctx context.Context, userID string) (map[string]float64, error) {
	rows, err := s.db.QueryContext(ctx, `
	SELECT tags, AVG(score) FROM question_records
	WHERE user_id=? AND created_at > ?
	GROUP BY tags`, userID, time.Now().AddDate(0, -3, 0))
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	result := make(map[string]float64)
	for rows.Next() {
		var tagsJSON string
		var avg float64
		rows.Scan(&tagsJSON, &avg)
		var tags []string
		json.Unmarshal([]byte(tagsJSON), &tags)
		for _, tag := range tags {
			tag = strings.TrimSpace(tag)
			if tag != "" {
				result[tag] = avg
			}
		}
	}
	return result, nil
}
