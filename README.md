# 大厂面试官 Agent

> 基于 AI 的大厂技术面试模拟平台，支持字节跳动、腾讯、阿里巴巴等 8 家公司的真实面试风格，含 AI Coding 专项面试模块。

---

## 功能介绍

### 面试模拟

- **8 家大厂风格**：字节跳动、腾讯、阿里巴巴、美团、京东、快手、小红书、拼多多
- **三轮面试结构**：
  - 一面：基础扎实性验证（技术基础 + 项目考察）
  - 二面：项目深挖（设计决策 + 踩坑经历 + 无限追问）
  - 三面：系统设计 + 业务落地（先聊业务再谈技术）
- **实时流式输出**：SSE 打字效果，接近真实面试官体验
- **每轮评估报告**：强项 / 薄弱点 / 建议 / 综合评级

### 简历与 JD 输入

- **PDF 简历自动解析**：纯 Go 实现，无需 API
- **截图 OCR**：Tesseract.js 浏览器端本地识别，支持中英文，无需 API
- **预设 JD 模板**：AI 编译器、LLM 推理加速、AI Infra、Go 后端、大模型预训练等
- **图片粘贴**：直接 Ctrl+V 粘贴 JD 截图，自动提取文字

### AI Coding 专项面试

- **14 道真实大厂场景题**：电商秒杀、分布式限流、RAG 系统、外卖调度、Meta 风格 AI 编程等
- **对话式面试**：AI 面试官主动提问 → 追问 → 实时点评
- **换一题**：随机切换题目

### 历史记录

- 面试记录本地持久化（JSON 文件），重启不丢失
- 随时回放历史对话
- 进行中的面试可一键「继续面试」

### 多 AI 供应商

在设置页面填入任意 API Key，无需改代码：

| 供应商 | 推荐指数 | 特点 |
|--------|---------|------|
| DeepSeek | ⭐⭐⭐⭐⭐ | 国内直连，便宜，效果好 |
| 智谱 GLM-4-Flash | ⭐⭐⭐⭐ | 有免费额度 |
| Claude (Anthropic) | ⭐⭐⭐⭐⭐ | 支持 PDF 原生识别 |
| GPT-4o (OpenAI) | ⭐⭐⭐⭐ | 综合能力强 |
| Qwen (阿里云) | ⭐⭐⭐ | 支持宝贝账号登录 |

---

## 快速开始

### 环境要求

- Go 1.21+（[安装指南](https://golang.org/doc/install)）

### 运行

```bash
# 1. 克隆项目
git clone https://github.com/chenyongzhi1119/interviewerAgent.git
cd interviewerAgent

# 2. 设置 AI 供应商（至少一个）
# 推荐 DeepSeek，国内可用，注册地址：https://platform.deepseek.com
export DEEPSEEK_API_KEY=sk-xxxxxxxx

# 也可以用其他供应商
# export ANTHROPIC_API_KEY=sk-ant-xxxxxxxx
# export OPENAI_API_KEY=sk-xxxxxxxx
# export GLM_API_KEY=xxxxxxxx
# export QWEN_API_KEY=xxxxxxxx

# 3. 启动
go run main.go

# 4. 打开浏览器
open http://localhost:8080
```

> **不想设置环境变量？** 启动后在页面右上角「设置」里填入 API Key，保存在浏览器本地，不上传服务器。

---

## 使用流程

```
1. 选择 AI 供应商 + 目标公司 + 面试轮次
2. 粘贴岗位 JD（支持文字 / 截图 / 预设模板）
3. 上传简历（PDF 自动解析 / 粘贴文字）
4. 点击「开始面试」→ AI 面试官主动提问
5. 用麦克风或键盘回答（Enter 发送）
6. 结束本轮 → 获取评估报告
7. 可选择继续下一轮面试
```

---

## 项目结构

```
interviewerAgent/
├── main.go                    # 入口，embed 前端文件
├── companies/                 # 公司面试风格模板（YAML）
│   ├── bytedance.yaml         # 字节跳动（无限追问式）
│   ├── tencent.yaml           # 腾讯（基础扎实优先）
│   ├── alibaba.yaml           # 阿里巴巴（业务价值 + 六脉神剑）
│   ├── meituan.yaml           # 美团（算法难度最高）
│   ├── jd.yaml                # 京东（电商物流实战）
│   ├── kuaishou.yaml          # 快手（短视频 + 推荐系统）
│   ├── xiaohongshu.yaml       # 小红书（内容社区 + 产品感）
│   └── pinduoduo.yaml         # 拼多多（高压高强度）
├── internal/
│   ├── agent/                 # 面试会话状态机
│   ├── llm/                   # AI 供应商抽象（Anthropic + OpenAI 兼容）
│   ├── extract/               # PDF 文字提取 + 图片 OCR
│   ├── model/                 # 数据模型
│   └── server/                # HTTP 路由 + SSE 流式
├── web/
│   ├── index.html             # 单页应用
│   ├── style.css              # 样式
│   ├── app.js                 # 前端逻辑
│   └── problems.js            # AI Coding 题库（14 道）
└── sessions/                  # 面试记录（本地存储，.gitignore 排除）
```

---

## 扩展：添加新公司

在 `companies/` 目录新建 `xxx.yaml`，重启服务自动加载：

```yaml
name: yourcompany
display_name: "公司名称"
role_description: |
  你是一名...面试官，...
rounds:
  1:
    title: "一面 · 标题"
    instructions: |
      本轮考察重点...
  2:
    title: "二面 · 标题"
    instructions: |
      ...
  3:
    title: "三面 · 标题"
    instructions: |
      ...
evaluation_rubric: |
  评估格式说明...
```

---

## 技术栈

| 层级 | 技术 |
|------|------|
| 后端 | Go + `net/http` 标准库，SSE 流式输出 |
| 前端 | 原生 HTML / CSS / JS，无框架依赖 |
| AI 供应商 | Anthropic SDK + OpenAI 兼容 API |
| PDF 解析 | `github.com/ledongthuc/pdf`（纯 Go） |
| 图片 OCR | Tesseract.js（浏览器端，无需 API） |
| 数据存储 | 本地 JSON 文件，无数据库依赖 |

---

## License

MIT License © 2025 chenyongzhi1119
