// ─── Problem Bank ─────────────────────────────────────────────────────────────

const ALGO_PROBLEMS = [
  {
    id: 1,
    title: '两数之和',
    leetcode_url: 'https://leetcode.cn/problems/two-sum/',
    difficulty: '简单',
    tags: ['数组', '哈希表'],
    companies: ['字节跳动', '腾讯', '阿里巴巴', '美团'],
    description: `给定一个整数数组 \`nums\` 和一个整数目标值 \`target\`，请你在该数组中找出**和为目标值 target** 的那两个整数，并返回它们的数组下标。

你可以假设每种输入只会对应一个答案，但是，数组中同一个元素不能使用两遍。

你可以按任意顺序返回答案。`,
    examples: [
      { input: 'nums = [2,7,11,15], target = 9', output: '[0,1]', explanation: 'nums[0] + nums[1] == 9，返回 [0, 1]' },
      { input: 'nums = [3,2,4], target = 6', output: '[1,2]', explanation: '' },
    ],
    constraints: ['2 ≤ nums.length ≤ 10⁴', '-10⁹ ≤ nums[i] ≤ 10⁹', '只会存在一个有效答案'],
    starter: {
      cpp: `class Solution {\npublic:\n    vector<int> twoSum(vector<int>& nums, int target) {\n        \n    }\n};`,
      python: `class Solution:\n    def twoSum(self, nums: List[int], target: int) -> List[int]:\n        `,
      go: `func twoSum(nums []int, target int) []int {\n    \n}`,
      java: `class Solution {\n    public int[] twoSum(int[] nums, int target) {\n        \n    }\n}`,
    }
  },
  {
    id: 2,
    title: '无重复字符的最长子串',
    leetcode_url: 'https://leetcode.cn/problems/longest-substring-without-repeating-characters/',
    difficulty: '中等',
    tags: ['哈希表', '字符串', '滑动窗口'],
    companies: ['字节跳动', '腾讯', '小红书'],
    description: `给定一个字符串 \`s\`，请你找出其中不含有重复字符的**最长子串**的长度。`,
    examples: [
      { input: 's = "abcabcbb"', output: '3', explanation: '最长子串为 "abc"，长度为 3' },
      { input: 's = "pwwkew"', output: '3', explanation: '最长子串为 "wke"，长度为 3' },
    ],
    constraints: ['0 ≤ s.length ≤ 5×10⁴', 's 由英文字母、数字、符号和空格组成'],
    starter: {
      cpp: `class Solution {\npublic:\n    int lengthOfLongestSubstring(string s) {\n        \n    }\n};`,
      python: `class Solution:\n    def lengthOfLongestSubstring(self, s: str) -> int:\n        `,
      go: `func lengthOfLongestSubstring(s string) int {\n    \n}`,
      java: `class Solution {\n    public int lengthOfLongestSubstring(String s) {\n        \n    }\n}`,
    }
  },
  {
    id: 3,
    title: 'LRU 缓存',
    leetcode_url: 'https://leetcode.cn/problems/lru-cache/',
    difficulty: '中等',
    tags: ['设计', '哈希表', '链表'],
    companies: ['字节跳动', '腾讯', '阿里巴巴', '美团', '快手'],
    description: `请你设计并实现一个满足 **LRU（最近最少使用）缓存** 约束的数据结构。

实现 \`LRUCache\` 类：
- \`LRUCache(int capacity)\` 以**正整数**作为容量 capacity 初始化 LRU 缓存
- \`int get(int key)\` 如果关键字 key 存在于缓存中，则返回关键字的值，否则返回 \`-1\`
- \`void put(int key, int value)\` 如果关键字 key 已经存在，则变更其数据值 value；如果不存在，则向缓存中插入该组 key-value。如果插入操作导致关键字数量超过 capacity，则应该逐出最久未使用的关键字

函数 **get** 和 **put** 必须以 **O(1)** 的平均时间复杂度运行。`,
    examples: [
      {
        input: 'LRUCache(2)\nput(1,1)\nput(2,2)\nget(1)\nput(3,3)\nget(2)\nput(4,4)\nget(1)\nget(3)\nget(4)',
        output: '[null,null,null,1,null,-1,null,1,3,4]',
        explanation: 'put(3,3) 时缓存满，淘汰最久未用的 key=2；put(4,4) 时淘汰 key=1',
      },
    ],
    constraints: ['1 ≤ capacity ≤ 3000', '0 ≤ key ≤ 10⁴', '0 ≤ value ≤ 10⁵', '最多调用 2×10⁵ 次 get 和 put'],
    starter: {
      cpp: `class LRUCache {\npublic:\n    LRUCache(int capacity) {\n        \n    }\n    int get(int key) {\n        \n    }\n    void put(int key, int value) {\n        \n    }\n};`,
      python: `class LRUCache:\n    def __init__(self, capacity: int):\n        \n    def get(self, key: int) -> int:\n        \n    def put(self, key: int, value: int) -> None:\n        `,
      go: `type LRUCache struct {\n    \n}\nfunc Constructor(capacity int) LRUCache {\n    \n}\nfunc (c *LRUCache) Get(key int) int {\n    \n}\nfunc (c *LRUCache) Put(key, value int) {\n    \n}`,
      java: `class LRUCache {\n    public LRUCache(int capacity) {\n        \n    }\n    public int get(int key) {\n        \n    }\n    public void put(int key, int value) {\n        \n    }\n}`,
    }
  },
  {
    id: 4,
    title: '二叉树的层序遍历',
    leetcode_url: 'https://leetcode.cn/problems/binary-tree-level-order-traversal/',
    difficulty: '中等',
    tags: ['树', '广度优先搜索', '二叉树'],
    companies: ['字节跳动', '腾讯', '美团'],
    description: `给你二叉树的根节点 \`root\`，返回其节点值的**层序遍历**结果（即逐层地，从左到右访问所有节点）。`,
    examples: [
      { input: 'root = [3,9,20,null,null,15,7]', output: '[[3],[9,20],[15,7]]', explanation: '' },
      { input: 'root = [1]', output: '[[1]]', explanation: '' },
    ],
    constraints: ['树中节点数目在范围 [0, 2000] 内', '-1000 ≤ Node.val ≤ 1000'],
    starter: {
      cpp: `class Solution {\npublic:\n    vector<vector<int>> levelOrder(TreeNode* root) {\n        \n    }\n};`,
      python: `class Solution:\n    def levelOrder(self, root: Optional[TreeNode]) -> List[List[int]]:\n        `,
      go: `func levelOrder(root *TreeNode) [][]int {\n    \n}`,
      java: `class Solution {\n    public List<List<Integer>> levelOrder(TreeNode root) {\n        \n    }\n}`,
    }
  },
  {
    id: 5,
    title: '接雨水',
    leetcode_url: 'https://leetcode.cn/problems/trapping-rain-water/',
    difficulty: '困难',
    tags: ['数组', '双指针', '动态规划', '单调栈'],
    companies: ['字节跳动', '阿里巴巴', '美团', '腾讯'],
    description: `给定 \`n\` 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水。`,
    examples: [
      { input: 'height = [0,1,0,2,1,0,1,3,2,1,2,1]', output: '6', explanation: '上面是由数组 [0,1,0,2,1,0,1,3,2,1,2,1] 表示的高度图，共接了 6 个单位的雨水' },
      { input: 'height = [4,2,0,3,2,5]', output: '9', explanation: '' },
    ],
    constraints: ['n == height.length', '1 ≤ n ≤ 2×10⁴', '0 ≤ height[i] ≤ 10⁵'],
    starter: {
      cpp: `class Solution {\npublic:\n    int trap(vector<int>& height) {\n        \n    }\n};`,
      python: `class Solution:\n    def trap(self, height: List[int]) -> int:\n        `,
      go: `func trap(height []int) int {\n    \n}`,
      java: `class Solution {\n    public int trap(int[] height) {\n        \n    }\n}`,
    }
  },
  {
    id: 6,
    title: '合并 K 个升序链表',
    leetcode_url: 'https://leetcode.cn/problems/merge-k-sorted-lists/',
    difficulty: '困难',
    tags: ['链表', '分治', '堆（优先队列）', '归并排序'],
    companies: ['字节跳动', '腾讯', '阿里巴巴'],
    description: `给你一个链表数组，每个链表都已经按升序排列。

请你将所有链表合并到一个升序链表中，返回合并后的链表。`,
    examples: [
      { input: 'lists = [[1,4,5],[1,3,4],[2,6]]', output: '[1,1,2,3,4,4,5,6]', explanation: '将三个链表合并：1→1→2→3→4→4→5→6' },
      { input: 'lists = []', output: '[]', explanation: '' },
    ],
    constraints: ['k == lists.length', '0 ≤ k ≤ 10⁴', '0 ≤ lists[i].length ≤ 500', '-10⁴ ≤ lists[i][j] ≤ 10⁴'],
    starter: {
      cpp: `class Solution {\npublic:\n    ListNode* mergeKLists(vector<ListNode*>& lists) {\n        \n    }\n};`,
      python: `class Solution:\n    def mergeKLists(self, lists: List[Optional[ListNode]]) -> Optional[ListNode]:\n        `,
      go: `func mergeKLists(lists []*ListNode) *ListNode {\n    \n}`,
      java: `class Solution {\n    public ListNode mergeKLists(ListNode[] lists) {\n        \n    }\n}`,
    }
  },
  {
    id: 7,
    title: '最小覆盖子串',
    leetcode_url: 'https://leetcode.cn/problems/minimum-window-substring/',
    difficulty: '困难',
    tags: ['哈希表', '字符串', '滑动窗口'],
    companies: ['字节跳动', '快手', '小红书'],
    description: `给你一个字符串 \`s\`、一个字符串 \`t\`。返回 \`s\` 中涵盖 \`t\` 所有字符的**最小子串**，如果 \`s\` 中不存在涵盖 \`t\` 所有字符的子串，则返回空字符串 \`""\`。

注意：对于 \`t\` 中重复字符，寻找的子串中该字符数量必须不少于 \`t\` 中该字符数量。`,
    examples: [
      { input: 's = "ADOBECODEBANC", t = "ABC"', output: '"BANC"', explanation: '最小覆盖子串 "BANC" 包含来自字符串 t 的 A、B 和 C' },
      { input: 's = "a", t = "aa"', output: '""', explanation: 't 中两个字符 a 均应包含在 s 的子串中，因此没有符合条件的子字符串' },
    ],
    constraints: ['1 ≤ s.length, t.length ≤ 10⁵', 's 和 t 由英文字母组成'],
    starter: {
      cpp: `class Solution {\npublic:\n    string minWindow(string s, string t) {\n        \n    }\n};`,
      python: `class Solution:\n    def minWindow(self, s: str, t: str) -> str:\n        `,
      go: `func minWindow(s string, t string) string {\n    \n}`,
      java: `class Solution {\n    public String minWindow(String s, String t) {\n        \n    }\n}`,
    }
  },
  {
    id: 8,
    title: '编辑距离',
    leetcode_url: 'https://leetcode.cn/problems/edit-distance/',
    difficulty: '中等',
    tags: ['字符串', '动态规划'],
    companies: ['字节跳动', '腾讯', '阿里巴巴'],
    description: `给你两个单词 \`word1\` 和 \`word2\`，请返回将 \`word1\` 转换成 \`word2\` 所使用的最少操作数。

你可以对一个单词进行如下三种操作：
- 插入一个字符
- 删除一个字符
- 替换一个字符`,
    examples: [
      { input: 'word1 = "horse", word2 = "ros"', output: '3', explanation: 'horse→rorse→rose→ros' },
      { input: 'word1 = "intention", word2 = "execution"', output: '5', explanation: '' },
    ],
    constraints: ['0 ≤ word1.length, word2.length ≤ 500', 'word1 和 word2 由小写英文字母组成'],
    starter: {
      cpp: `class Solution {\npublic:\n    int minDistance(string word1, string word2) {\n        \n    }\n};`,
      python: `class Solution:\n    def minDistance(self, word1: str, word2: str) -> int:\n        `,
      go: `func minDistance(word1 string, word2 string) int {\n    \n}`,
      java: `class Solution {\n    public int minDistance(String word1, String word2) {\n        \n    }\n}`,
    }
  },
];

// ─── AI Coding Problems ────────────────────────────────────────────────────────

const AI_CODING_PROBLEMS = [
  {
    id: 1,
    title: '电商购物车并发服务',
    company: '淘天集团',
    scenario: '双十一大促场景',
    difficulty: '中等',
    tags: ['并发', 'Redis', 'Go', '幂等性'],
    background: `双十一大促期间，购物车服务承受每秒数万次并发请求。你需要借助 AI 工具（Copilot / ChatGPT / Claude 等）完成核心模块的设计与实现。`,
    requirements: [
      '实现添加/删除商品、查询购物车接口（RESTful）',
      '基于 Redis 实现库存预扣减，防止超卖',
      '接口幂等性保证（重复请求不重复扣减）',
      '支持用户会话过期自动清理',
    ],
    tech_stack: 'Go + Redis + MySQL',
    evaluation_points: [
      '**Prompt 设计**：给 AI 的需求描述是否清晰完整，上下文约束是否合理',
      '**架构决策**：技术选型（如分布式锁方案）的合理性，能否清晰解释',
      '**代码审查**：对 AI 生成代码的审查能力，是否识别出潜在 bug',
      '**边界处理**：是否考虑了网络超时、Redis 故障降级等异常场景',
    ],
    starter_prompt: `请帮我实现一个 Go 语言的购物车后端服务，要求：
1. 使用 Redis 存储购物车数据，key 格式为 cart:{userId}
2. 实现 AddItem、RemoveItem、GetCart 三个接口
3. 添加商品时需要校验库存并预扣减（Redis DECR 原子操作）
4. 使用 Redis 分布式锁保证并发安全

请给出完整的 Go 代码实现，包括数据结构定义和接口实现。`
  },
  {
    id: 2,
    title: 'Feed 流实时推荐接口',
    company: '字节跳动',
    scenario: '抖音 / 今日头条场景',
    difficulty: '中等',
    tags: ['推荐系统', 'API 设计', '缓存', 'Go'],
    background: `你负责字节跳动 Feed 流推荐系统的后端接口开发。需要使用 AI Coding 工具完成拉取推荐内容、记录用户行为两个核心接口的实现。`,
    requirements: [
      '设计 /feed/recommend 接口，支持分页、去重（已看过的内容不重复推送）',
      '设计 /feed/action 接口，记录用户点赞/播放/分享行为',
      '使用 Redis 缓存推荐列表，TTL 5 分钟',
      '行为数据异步写入 Kafka，不影响主链路延迟',
    ],
    tech_stack: 'Go + Redis + Kafka',
    evaluation_points: [
      '**接口设计**：Request/Response 结构是否合理，字段命名是否规范',
      '**异步设计**：Kafka 生产者的错误处理和重试机制',
      '**缓存策略**：缓存击穿/穿透的防护设计',
      '**Prompt 质量**：AI 生成代码的质量与你的 Prompt 设计的关系',
    ],
    starter_prompt: `帮我实现两个 Go HTTP 接口：
1. GET /feed/recommend?userId=xxx&page=1&size=10
   - 从 Redis 获取推荐内容列表（key: feed:recommend:{userId}）
   - 过滤掉用户已浏览的内容（从 Redis Set: feed:seen:{userId} 判断）
   - 返回 JSON：{code, data: [{id, title, type, coverUrl}], hasMore}

2. POST /feed/action
   - Body: {userId, contentId, action: "like|play|share"}
   - 异步发送到 Kafka topic: user-behavior
   - 将 contentId 加入用户已浏览集合

请使用 gin 框架，包含完整的错误处理。`
  },
  {
    id: 3,
    title: '分布式限流器',
    company: '通用（大厂高频）',
    scenario: '微服务网关场景',
    difficulty: '困难',
    tags: ['限流', 'Redis', '分布式', '算法'],
    background: `微服务架构中，网关层需要对下游服务进行 QPS 限流，防止流量突增压垮服务。你需要使用 AI 工具实现一个支持多种限流算法的分布式限流器。`,
    requirements: [
      '实现令牌桶（Token Bucket）算法：平滑限流，允许突发流量',
      '实现滑动窗口（Sliding Window）算法：精确统计请求数',
      '基于 Redis + Lua 脚本实现原子操作，保证分布式场景正确性',
      '支持按 IP、用户 ID、接口路径 等维度限流',
    ],
    tech_stack: 'Go + Redis + Lua',
    evaluation_points: [
      '**算法理解**：能否解释令牌桶和滑动窗口的区别及适用场景',
      '**原子性保证**：Lua 脚本设计是否正确，是否理解 Redis 单线程的意义',
      '**抽象设计**：限流器接口抽象是否合理，方便后续扩展新算法',
      '**AI 工具使用**：能否通过迭代 Prompt 让 AI 修复 Lua 脚本中的 bug',
    ],
    starter_prompt: `实现一个基于 Redis + Lua 的分布式令牌桶限流器（Go 语言）：

接口定义：
type Limiter interface {
    Allow(ctx context.Context, key string) (bool, error)
}

要求：
1. 构造函数接收 rate（每秒令牌数）和 burst（桶容量）
2. Allow 方法通过 Lua 脚本原子地执行：获取当前令牌数 → 计算补充令牌 → 判断是否允许
3. Redis Key 格式：limiter:{key}，存储 {tokens, lastRefill} 两个字段
4. 使用 go-redis/v9 客户端

请给出完整实现，包括 Lua 脚本。`
  },
  {
    id: 4,
    title: '代码质量自动 Review Agent',
    company: 'AI 应用专项',
    scenario: 'AI Coding 综合能力考察',
    difficulty: '中等',
    tags: ['AI Agent', 'LLM', 'Prompt Engineering', 'Go'],
    background: `这是一道 AI 应用开发题，考察你构建 AI 工具的能力。你需要设计一个 Go 程序，调用 LLM API 对提交的代码进行自动 Review，并给出结构化的反馈报告。`,
    requirements: [
      '接受 stdin 输入代码，调用 LLM API（OpenAI 格式）进行 Review',
      '输出结构化报告：{bugs, performance, style, security, score}',
      '支持流式输出（SSE），实时展示 Review 过程',
      '系统 Prompt 设计：角色、输出格式、评分标准',
    ],
    tech_stack: 'Go + Any LLM API',
    evaluation_points: [
      '**Prompt Engineering**：System Prompt 设计是否使 LLM 输出稳定的 JSON 结构',
      '**流式处理**：SSE 解析和转发的正确性',
      '**错误处理**：LLM 返回非 JSON 时的降级处理',
      '**扩展性**：如何支持多模型（GPT-4 / Claude / DeepSeek）切换',
    ],
    starter_prompt: `帮我实现一个 Go 命令行工具，功能是调用 LLM API 对代码进行 Review：

1. 读取 stdin 的代码内容
2. 构造请求发送到 OpenAI Chat API（支持流式输出）
3. System Prompt 要求 LLM 以 JSON 格式返回：
   {
     "bugs": ["bug描述1", "bug描述2"],
     "performance": ["性能问题1"],
     "style": ["代码风格问题1"],
     "security": ["安全风险1"],
     "score": 85,
     "summary": "总体评价"
   }
4. 流式打印 LLM 的输出，最后解析 JSON 美化输出

请实现完整的 Go 代码，使用标准库 net/http 调用 API。`
  },

  // ── 国内大厂场景 ─────────────────────────────────────────────
  {
    id: 5,
    title: '优惠券发放与核销系统',
    company: '美团 / 饿了么',
    scenario: '大促活动场景',
    difficulty: '中等',
    tags: ['幂等性', 'Redis', '分布式事务', 'Go'],
    background: `双十一大促期间，平台需要向千万级用户发放各类优惠券（满减、折扣、免运费），并在下单时完成核销。高并发下需防止重复发券、超发和重复核销。`,
    requirements: [
      '发券接口：用户领取优惠券，同一用户同一券种限领一张',
      '核销接口：下单时核销优惠券，保证幂等性（相同订单 ID 多次调用不重复核销）',
      '库存管理：Redis 原子操作控制券总量，防止超发',
      '过期处理：券有效期内可用，过期自动失效',
    ],
    tech_stack: 'Go + Redis + MySQL',
    evaluation_points: [
      '**幂等性设计**：发券和核销的幂等方案（Lua 脚本 / 唯一索引 / 状态机），能否清晰解释选型理由',
      '**并发控制**：高并发下的超发防护，是否理解 Redis SETNX / DECR 的原子性边界',
      '**Prompt 精确度**：给 AI 的 Prompt 能否清晰描述幂等需求和并发约束',
      '**错误处理**：网络超时、Redis 故障时的降级策略',
    ],
    starter_prompt: `帮我实现优惠券发放接口（Go + Redis + MySQL）：

接口：POST /api/coupon/grant
Request: {userId, couponTemplateId}
Response: {code, couponId, expireAt}

业务规则：
1. 检查券模板库存（Redis key: coupon:stock:{templateId}，初始值=总库存）
2. 检查用户是否已领取（Redis Set: coupon:granted:{templateId} 存放已领 userId）
3. 如果未领且有库存：MULTI/EXEC 原子扣减库存 + 记录用户
4. 写入 MySQL coupon 表：(id, user_id, template_id, status=unused, expire_at)
5. 返回优惠券 ID

请用 Go + go-redis/v9 实现，包含完整错误处理。`,
  },
  {
    id: 6,
    title: '实时消息推送服务',
    company: '腾讯 / 钉钉',
    scenario: 'IM 消息系统',
    difficulty: '中等',
    tags: ['WebSocket', 'Go', '消息队列', '在线状态'],
    background: `IM 系统需要将消息实时推送给在线用户。你需要借助 AI 工具实现一个轻量级的 WebSocket 消息推送服务，支持单聊和群聊消息的实时投递。`,
    requirements: [
      '维护用户在线连接表（userId → WebSocket连接）',
      '接收来自 Kafka 的消息事件，推送给目标用户',
      '用户断线重连时，推送离线期间的未读消息',
      '心跳检测：60s 无消息自动断开，客户端每30s ping一次',
    ],
    tech_stack: 'Go + WebSocket + Kafka + Redis',
    evaluation_points: [
      '**连接管理**：并发读写 map 的线程安全处理（sync.Map / RWMutex），是否意识到这个坑',
      '**离线消息**：离线消息的存储和拉取策略（Redis List / 推拉结合）',
      '**Prompt 质量**：WebSocket 生命周期管理的 Prompt 描述是否准确',
      '**资源泄漏**：连接关闭时的清理逻辑，goroutine 是否会泄漏',
    ],
    starter_prompt: `用 Go 实现一个 WebSocket 消息推送服务器：

要求：
1. 用 gorilla/websocket 处理 WebSocket 连接
2. 全局连接管理器 ConnManager：sync.Map 存储 userId → *websocket.Conn
3. 处理函数 handleConn(userId, conn)：
   - 启动 goroutine 读取客户端消息（ping/pong）
   - 60s 超时自动关闭
   - 断开时从 ConnManager 删除
4. 推送函数 Push(userId, msg)：找到连接并发送消息，用户离线则写入 Redis List (offline:msgs:{userId})
5. HTTP 路由：GET /ws?token=xxx 升级为 WebSocket

请给出完整实现，包含连接管理器和推送函数。`,
  },
  {
    id: 7,
    title: '外卖骑手智能调度系统',
    company: '美团外卖',
    scenario: '本地生活高频场景',
    difficulty: '困难',
    tags: ['地理位置', '调度算法', 'Go', 'Redis GeoHash'],
    background: `外卖平台每天处理数千万订单，核心挑战是实时将新订单分配给最合适的骑手。你需要借助 AI 工具设计并实现订单分配模块的核心逻辑。`,
    requirements: [
      '骑手签到/签退接口：更新骑手位置到 Redis GeoHash',
      '订单分配接口：接单时在 3km 内找最近的空闲骑手',
      '骑手状态管理：空闲/取餐中/配送中 三种状态',
      '超时未接单：120s 内无骑手接单，升级为紧急订单扩大搜索半径',
    ],
    tech_stack: 'Go + Redis GeoSearch + MySQL',
    evaluation_points: [
      '**地理位置方案**：Redis GEOSEARCH 的使用，能否解释 GeoHash 精度和半径的关系',
      '**状态并发**：骑手状态更新的竞争条件（两个订单同时分配给同一骑手），如何解决',
      '**超时机制**：Redis 延迟队列（ZSET）实现超时重分配的方案',
      '**AI Prompt 层次**：先问设计方案再让 AI 写代码，还是一步到位',
    ],
    starter_prompt: `实现外卖骑手调度模块（Go + Redis）：

1. 骑手位置更新：
   POST /rider/location {riderId, lat, lng, status}
   用 Redis GEOADD 更新位置（key: riders:geo）

2. 订单分配：
   POST /order/assign {orderId, restaurantLat, restaurantLng}
   用 GEOSEARCH riders:geo FROMLONLAT {lng} {lat} BYRADIUS 3 km ASC COUNT 10
   过滤 status=idle 的骑手，取最近的一个
   原子性更新骑手状态为 busy（用 Redis SET NX 实现乐观锁）

3. 超时处理：
   分配成功后，向 Redis ZSET (key: order:timeout) 加入 {orderId, score=now+120}
   后台 goroutine 轮询检查超时订单

请实现完整 Go 代码。`,
  },
  {
    id: 8,
    title: '智能搜索补全服务',
    company: '百度 / 阿里',
    scenario: '搜索框实时补全',
    difficulty: '中等',
    tags: ['Trie', 'Redis', 'Go', '热词统计'],
    background: `搜索框的实时补全（autocomplete）是高频用户交互场景，需要在 50ms 内返回 Top 5 候选词。你需要用 AI 工具实现一个支持中英文混合、按热度排序的搜索补全服务。`,
    requirements: [
      '实时补全接口：输入前缀，返回 Top 5 热词（延迟 < 50ms）',
      '热词更新：用户搜索后异步更新该词的搜索次数',
      '使用 Redis Sorted Set 存储词频，前缀用 ZRANGEBYLEX 范围查询',
      '支持拼音首字母补全（如输入"gjy"补全"工具栏"）',
    ],
    tech_stack: 'Go + Redis + 异步消息队列',
    evaluation_points: [
      '**Redis 数据结构选型**：ZSET 做前缀搜索的思路（成员按字典序，score=0，用范围查询），能否独立想到',
      '**性能意识**：为什么不用 MySQL LIKE 查询，能否量化性能差距',
      '**拼音方案**：如何让 AI 帮你生成拼音索引，Prompt 怎么写',
      '**并发写入**：高并发热词更新的批量合并策略（避免频繁写 Redis）',
    ],
    starter_prompt: `实现搜索词自动补全服务（Go + Redis）：

数据结构：
- Redis ZSET key: search:words，所有词的 score=0（利用字典序），用 ZRANGEBYLEX 做前缀匹配
- Redis ZSET key: search:hot，词→搜索次数，用于热度排序

补全接口 GET /search/suggest?q=前缀：
1. ZRANGEBYLEX search:words [q (q\xff LIMIT 0 100 → 得到前缀匹配的词列表
2. 用这些词在 search:hot 里批量获取热度（ZSCORE 或 ZMSCORE）
3. 按热度降序取 Top 5 返回

热词更新 POST /search/record {word}：
- 异步写入（channel buffer），批量合并后 ZINCRBY search:hot 1 word
- 如果是新词，同时 ZADD search:words 0 word

请实现完整 Go 代码。`,
  },
  {
    id: 9,
    title: '日志异常检测 Agent',
    company: '字节跳动 基础架构',
    scenario: 'SRE / 可观测性场景',
    difficulty: '中等',
    tags: ['AI Agent', 'LLM', 'Go', '日志分析'],
    background: `SRE 团队每天面对海量服务日志。你需要借助 AI 工具构建一个日志异常检测 Agent：接收日志流，自动识别异常模式，生成可读的告警摘要。`,
    requirements: [
      '接收 stdin 输入的日志文本（JSON 格式，多行）',
      '调用 LLM 分析日志，识别 ERROR/异常堆栈/超时等模式',
      '输出结构化告警：{level, pattern, affected_services, suggestion, count}',
      '支持批量分析（每100条日志一批），避免 token 超限',
    ],
    tech_stack: 'Go + Any LLM API',
    evaluation_points: [
      '**System Prompt 设计**：如何让 LLM 稳定输出 JSON 格式，few-shot 样例的选取',
      '**分批策略**：日志分批的边界处理（不能截断一个异常栈的上下文）',
      '**Prompt 迭代**：当 LLM 输出格式不稳定时，如何迭代 Prompt 修复',
      '**成本意识**：如何在保证质量的前提下减少 token 消耗（预过滤、摘要压缩）',
    ],
    starter_prompt: `构建一个日志异常检测工具（Go）：

输入：从 stdin 读取 JSON 日志，每行一条：
{"timestamp":"2024-01-01T10:00:00Z","level":"ERROR","service":"payment","msg":"connection timeout","trace_id":"abc123"}

处理逻辑：
1. 每积累 50 条日志（或遇到 EOF）触发一次 LLM 分析
2. System Prompt：你是 SRE 工程师，分析日志找出异常模式，以 JSON 格式输出
3. 要求 LLM 输出：{"anomalies":[{"pattern":"connection timeout","services":["payment"],"count":3,"severity":"HIGH","suggestion":"检查数据库连接池配置"}]}
4. 解析输出并打印告警报告

请实现完整 Go 代码，包含 LLM 调用（OpenAI 格式）和日志批处理。`,
  },
  {
    id: 10,
    title: 'RAG 知识库问答系统',
    company: 'AI 应用专项',
    scenario: '企业内部知识库',
    difficulty: '中等',
    tags: ['RAG', 'Embedding', 'Vector DB', 'Go/Python'],
    background: `企业需要将内部文档（PDF、Markdown、Wiki）接入 LLM，让员工用自然语言查询。你需要借助 AI 工具实现 RAG（检索增强生成）系统的核心流程。`,
    requirements: [
      '文档摄入：读取文档，按段落切分，调用 Embedding API 向量化',
      '向量存储：将向量和原文存入向量数据库（可用内存 map 简化）',
      '查询接口：输入问题 → Embedding → 向量相似度检索 → 拼装 Prompt → LLM 生成答案',
      '引用溯源：返回答案时附上来源文档片段',
    ],
    tech_stack: 'Go 或 Python + OpenAI/类 OpenAI Embedding API',
    evaluation_points: [
      '**切分策略**：chunk_size 和 chunk_overlap 的选取，能否解释对召回质量的影响',
      '**相似度计算**：余弦相似度 vs 点积，为什么 Embedding 后用余弦',
      '**Prompt 工程**：RAG Prompt 的结构（Context + Question + 输出格式约束）',
      '**召回质量**：如何让 AI 帮你写测试用例验证召回效果',
    ],
    starter_prompt: `实现一个简单的 RAG 知识库问答系统（Python）：

1. 文档摄入函数 ingest(text, doc_id)：
   - 按 500 字分块，相邻块 50 字重叠
   - 调用 OpenAI text-embedding-3-small 向量化
   - 存入内存 dict: {doc_id: [(chunk_text, embedding_vector)]}

2. 检索函数 retrieve(query, top_k=3)：
   - 向量化 query
   - 对所有 chunk 计算余弦相似度
   - 返回 top_k 最相似的 chunk

3. 问答函数 ask(question)：
   - 检索 top 3 相关 chunk
   - 构造 Prompt："根据以下资料回答问题：\n{chunks}\n\n问题：{question}"
   - 调用 GPT-4o 生成答案，附上来源 chunk

请实现完整 Python 代码。`,
  },
  {
    id: 11,
    title: 'Meta 风格：迷宫 BFS 导航（AI 辅助）',
    company: 'Meta（国际大厂参考题型）',
    scenario: '2025 年 Meta AI 编程轮真实题型',
    difficulty: '困难',
    tags: ['BFS', '图算法', 'AI Coding', '代码扩展'],
    background: `Meta 2025 年引入的 AI 编程轮：给你一段已有的迷宫类代码骨架，需要借助 AI 工具在 60 分钟内将功能完整实现并通过测试。这道题考察你与 AI 协作、审查 AI 代码并修复 bug 的能力。

迷宫规则：
- 二维网格，.=可通行，#=墙，S=起点，E=终点
- 部分格子有方向门（→←↑↓），只能从指定方向进入
- 找从 S 到 E 的最短路径，返回步数和路径坐标`,
    requirements: [
      '实现 BFS 最短路径搜索，处理方向门约束',
      '复用 AI 生成的格子解析代码，但需检查边界处理 bug',
      '输出最短步数和完整路径，无路径时返回 -1',
      '编写至少 3 个测试用例（正常/无解/有方向门）',
    ],
    tech_stack: 'Python 或 Go（任选）',
    evaluation_points: [
      '**Prompt 分层**：先让 AI 生成类骨架，再让它填充具体方法，最后让它写测试——分层 Prompt 策略',
      '**Bug 识别**：AI 生成的边界检查代码是否有 off-by-one 错误，你能否快速发现',
      '**测试驱动**：用 AI 生成测试用例，再用测试发现代码问题的工作流',
      '**时间管理**：60 分钟内如何分配：需求理解 → Prompt 设计 → 代码生成 → 审查修复 → 测试',
    ],
    starter_prompt: `我在做一道迷宫 BFS 题，给你这个 Python 骨架，请填充实现：

\`\`\`python
class Maze:
    def __init__(self, grid: list[str]):
        self.grid = grid
        self.rows = len(grid)
        self.cols = len(grid[0]) if grid else 0
        # 方向门：某格只能从特定方向进入
        # doors[(r,c)] = {'→','←','↑','↓'} 表示该格允许从哪些方向进入
        self.doors = {}
        self._parse_doors()

    def _parse_doors(self):
        # TODO: 解析 grid 中的方向门字符，填充 self.doors
        pass

    def find_start_end(self):
        # TODO: 返回 (start_pos, end_pos)
        pass

    def bfs(self) -> tuple[int, list[tuple]]:
        # TODO: BFS 搜索，返回 (最短步数, 路径坐标列表)
        # 无解返回 (-1, [])
        pass
\`\`\`

请填充三个方法，并写 3 个测试用例覆盖：
1. 普通迷宫（有解）
2. 无解迷宫
3. 含方向门的迷宫`,
  },
  {
    id: 12,
    title: '高并发抢购系统压测与优化',
    company: '京东 / 拼多多',
    scenario: '618 大促秒杀场景',
    difficulty: '困难',
    tags: ['压测', '性能优化', 'Go', 'Redis', 'AI辅助分析'],
    background: `618 大促前夕，抢购系统在压测中发现 QPS 上限只有 2000，远低于预期的 50000。你需要借助 AI 工具分析性能瓶颈并给出优化方案。系统已有基础代码，你需要识别问题并用 AI 辅助重构。`,
    requirements: [
      '给 AI 提供压测日志和慢查询日志，让它分析瓶颈所在',
      '优化方案：Redis 缓存库存（减少 DB 查询）、异步下单（MQ 削峰）、本地缓存（减少 Redis 请求）',
      '实现令牌桶限流保护接口（防止单用户刷请求）',
      '优化后压测对比：目标 QPS ≥ 30000，P99 < 100ms',
    ],
    tech_stack: 'Go + Redis + Kafka',
    evaluation_points: [
      '**问题诊断**：能否把压测日志清晰地描述给 AI，让它识别出是 DB 锁竞争还是 Redis 热点',
      '**方案权衡**：超卖防护 vs 性能，强一致 vs 最终一致，能否清晰解释选择',
      '**分层优化**：本地缓存 → Redis → DB 的三级缓存方案，Prompt 如何分步实现',
      '**回滚方案**：如果优化后出现数据不一致，如何用 AI 快速生成数据修复脚本',
    ],
    starter_prompt: `我的秒杀接口性能不达标，以下是当前实现（Go），请分析瓶颈并给出优化版本：

\`\`\`go
// 当前实现：每次请求直接查询 MySQL
func (s *SeckillService) BuyItem(ctx context.Context, userId, itemId int64) error {
    tx := s.db.Begin()
    var item Item
    // 悲观锁查库存
    tx.Set("gorm:query_option", "FOR UPDATE").First(&item, itemId)
    if item.Stock <= 0 {
        tx.Rollback()
        return errors.New("已售罄")
    }
    tx.Model(&item).Update("stock", item.Stock-1)
    order := Order{UserId: userId, ItemId: itemId, Status: "paid"}
    tx.Create(&order)
    return tx.Commit().Error
}
\`\`\`

请：
1. 分析这段代码的性能瓶颈（重点：锁竞争、DB 写入压力）
2. 给出优化版本：Redis 预扣库存 + 异步落库（Kafka）
3. 添加令牌桶限流（每秒最多处理 50000 请求）`,
  },
  {
    id: 13,
    title: '多模态内容审核 Pipeline',
    company: '快手 / 抖音',
    scenario: '短视频内容安全',
    difficulty: '中等',
    tags: ['AI Pipeline', 'Python', '多模态', '异步处理'],
    background: `短视频平台每天上传数百万条视频，需要对图文内容进行实时审核。你需要借助 AI 工具设计并实现一个多阶段审核 Pipeline，结合规则引擎和大模型。`,
    requirements: [
      '阶段1：关键词过滤（规则引擎，< 1ms）',
      '阶段2：图像/封面帧检测（调用视觉 API，< 500ms）',
      '阶段3：LLM 深度审核（对前两阶段标记可疑的内容，< 3s）',
      '输出：{status: pass/reject/review, reason, confidence, stage}',
    ],
    tech_stack: 'Python + OpenAI Vision API + 异步并发',
    evaluation_points: [
      '**Pipeline 设计**：三阶段过滤的数量漏斗设计，如何估算每阶段的过滤率',
      '**并发处理**：asyncio 实现批量并发调用视觉 API，Prompt 如何描述异步需求',
      '**成本控制**：只有前两阶段过不了才调 LLM，节省约 90% 的 LLM 成本——能否向 AI 解释这个设计',
      '**边界情况**：视觉 API 超时、LLM 返回格式错误时的 fallback 策略',
    ],
    starter_prompt: `实现一个三阶段内容审核 Pipeline（Python + asyncio）：

class ContentModerator:
    def __init__(self, openai_client):
        self.client = openai_client
        self.keyword_blacklist = ["暴力", "违禁", ...]

    async def stage1_keyword_filter(self, text: str) -> tuple[bool, str]:
        """规则过滤，返回 (通过, 原因)"""
        pass

    async def stage2_vision_check(self, image_url: str) -> tuple[bool, float]:
        """调用 GPT-4o Vision 检测图像，返回 (通过, 置信度)"""
        pass

    async def stage3_llm_deep_review(self, text: str, image_url: str) -> dict:
        """LLM 深度审核，返回 {pass: bool, reason: str}"""
        pass

    async def moderate(self, text: str, image_url: str) -> dict:
        """串联三阶段，只有前阶段触发才进入下一阶段"""
        pass

请填充实现，并给出 5 个测试用例（正常/违规文字/违规图像/两者都违规/API超时）。`,
  },
  {
    id: 14,
    title: 'SQL 慢查询自动优化助手',
    company: '通用（DBA 工具）',
    scenario: '数据库性能优化',
    difficulty: '中等',
    tags: ['SQL优化', 'LLM', 'Go', 'MySQL'],
    background: `数据库慢查询是后端性能问题的重要来源。你需要借助 AI 工具构建一个慢查询分析助手：输入 SQL 和 EXPLAIN 结果，自动给出优化建议和重写后的 SQL。`,
    requirements: [
      '解析 MySQL EXPLAIN 输出，识别全表扫描、索引失效、行数过多等问题',
      '调用 LLM 分析原因并给出优化建议（加索引/改写 SQL/分页优化）',
      '自动生成优化后的 SQL，并解释每处改动的原因',
      '批量处理：支持从慢查询日志文件读取多条 SQL',
    ],
    tech_stack: 'Go + MySQL + LLM API',
    evaluation_points: [
      '**Prompt 结构**：如何把 SQL + EXPLAIN 结果有效组织成 LLM 可理解的格式',
      '**输出约束**：让 LLM 同时输出"问题分析"+"优化后 SQL"+"改动说明"，格式稳定性如何保证',
      '**知识边界**：LLM 给出的索引建议是否合理，你如何验证（EXPLAIN 对比）',
      '**Prompt 进阶**：能否让 AI 模拟 DBA 思维，考虑索引对写性能的影响',
    ],
    starter_prompt: `构建 SQL 慢查询优化助手（Go）：

输入结构：
{
  "sql": "SELECT * FROM orders WHERE user_id=123 AND status='pending' ORDER BY created_at DESC",
  "explain_result": "| id | select_type | table  | type | possible_keys | key  | rows   | Extra       |\n| 1  | SIMPLE      | orders | ALL  | NULL          | NULL | 500000 | Using where |"
}

处理流程：
1. 构造 Prompt，包含：SQL 原文 + EXPLAIN 结果 + 要求（JSON 格式输出）
2. LLM 返回：{"issues": ["全表扫描500万行"], "optimized_sql": "...", "recommendations": ["在(user_id, status)上建联合索引"], "explanation": "..."}
3. 格式化输出优化报告

请实现完整 Go 代码，支持从命令行参数传入 JSON 文件路径，批量处理多条慢查询。`,
  },
];
