# 大厂面试官 Agent · InterviewPro

一个基于 AI 的大厂技术面试模拟系统，支持字节跳动、腾讯、阿里巴巴等 8 家公司的三轮面试风格，以及 AI Coding 专项面试。

## 功能特性

- **多公司面试风格**：字节跳动、腾讯、阿里巴巴、美团、京东、快手、小红书、拼多多
- **三轮面试结构**：一面（基础扎实性）、二面（项目深挖）、三面（系统设计 + 业务落地）
- **AI Coding 面试**：14 道真实大厂场景题，对话式面试，考察 AI 工具使用能力
- **多 AI 供应商**：支持 Claude、GPT-4o、DeepSeek、GLM、Qwen，在设置页面配置 API Key
- **简历智能解析**：PDF 自动提取文字，截图 OCR（Tesseract.js 本地识别，无需 API）
- **历史记录持久化**：面试记录本地保存，支持随时回放和继续未完成的面试
- **流式输出**：SSE 实时打字效果，接近真实面试官体验

## 快速开始

### 环境要求

- Go 1.21+

### 安装运行

```bash
git clone https://github.com/chenyongzhi1119/interviewerAgent.git
cd interviewerAgent

# 设置至少一个 AI 供应商（也可在页面设置中配置）
export DEEPSEEK_API_KEY=your_key_here   # 推荐：国内可用，便宜
# export ANTHROPIC_API_KEY=your_key_here  # Claude
# export OPENAI_API_KEY=your_key_here     # GPT-4o

go run main.go
```

打开浏览器访问 `http://localhost:8080`

### 支持的 AI 供应商

| 供应商 | 环境变量 | 默认模型 | 注册地址 |
|--------|---------|---------|---------|
| DeepSeek | `DEEPSEEK_API_KEY` | deepseek-chat | [platform.deepseek.com](https://platform.deepseek.com) |
| 智谱 GLM | `GLM_API_KEY` | glm-4-flash | [open.bigmodel.cn](https://open.bigmodel.cn) |
| Claude | `ANTHROPIC_API_KEY` | claude-sonnet-4-6 | [console.anthropic.com](https://console.anthropic.com) |
| OpenAI | `OPENAI_API_KEY` | gpt-4o | [platform.openai.com](https://platform.openai.com) |
| 阿里 Qwen | `QWEN_API_KEY` | qwen-plus | [dashscope.aliyuncs.com](https://dashscope.aliyuncs.com) |

也可以不设置环境变量，直接在页面右上角「设置」中填入 API Key（保存在浏览器 localStorage，不上传服务器）。

## 技术栈

- **后端**：Go + `net/http` 标准库，SSE 流式输出
- **前端**：原生 HTML/CSS/JS，无框架依赖
- **AI**：Anthropic SDK（Claude）+ OpenAI 兼容 API（DeepSeek/GLM/Qwen）
- **OCR**：Tesseract.js（浏览器端，无需 API）
- **PDF 解析**：`github.com/ledongthuc/pdf`（纯 Go，无需外部工具）

## 项目结构

```
interviewerAgent/
├── main.go                    # 入口，embed 前端静态文件
├── companies/                 # 公司面试风格模板（YAML）
│   ├── bytedance.yaml
│   ├── tencent.yaml
│   └── ...（8 家公司）
├── internal/
│   ├── agent/                 # 面试会话状态机
│   ├── llm/                   # AI 供应商抽象层
│   ├── extract/               # PDF 和图片文字提取
│   ├── model/                 # 数据模型
│   └── server/                # HTTP 路由层
├── web/                       # 前端文件（embed 进二进制）
│   ├── index.html
│   ├── style.css
│   ├── app.js
│   └── problems.js            # AI Coding 题库（14 道）
└── sessions/                  # 面试记录（本地存储，.gitignore 排除）
```

## 扩展：添加新公司

在 `companies/` 目录新建 `xxx.yaml` 文件，参考现有模板格式，重启服务即自动加载：

```yaml
name: yourcompany
display_name: "公司名称"
role_description: |
  面试官角色描述...
rounds:
  1:
    title: "一面"
    instructions: |
      本轮考察重点...
evaluation_rubric: |
  评估标准...
```

## License

MIT
