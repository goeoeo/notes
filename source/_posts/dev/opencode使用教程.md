# OpenCode MCP 服务器功能介绍

本文档介绍 opencode 中已接入的 MCP（Model Context Protocol）服务器及其各自的作用。MCP 是开放协议，让 AI 助手能够对接外部工具、代码、文档和知识库。

## MCP 列表

| 编号 | 名称 | 状态 | 来源 |
| --- | --- | --- | --- |
| 1 | `ast_grep` | Connected | 内置 |
| 2 | `codebase-memory-mcp` | Connected | 用户安装 |
| 3 | `confluence` | Connected | 用户安装 |
| 4 | `context7` | Connected | 内置 |
| 5 | `grep_app` | Connected | 内置 |
| 6 | `lsp` | Connected | 内置 |
| 7 | `vision-ocr-mcp` | Connected | 用户安装 |
| 8 | `websearch` | Connected | 内置 |
| 9 | `wechat-devtools` | Disabled | 用户安装 |

---

## 1. ast_grep — 代码结构化搜索与改写

**是什么**：基于 AST（抽象语法树）的代码搜索和批量改写工具，不是简单的文本正则匹配。

**典型用途**：
- 按语法结构调整代码模式，例如把 `console.log($MSG)` 批量改成 `logger.info($MSG)`；
- 支持 25 种编程语言，元变量（`$VAR` / `$$$`）匹配整个语法节点而非子字符串；
- 干跑（dry-run）模式可先预览改动再应用，适合安全重构。

---

## 2. codebase-memory-mcp — 代码库知识图谱

**是什么**：为代码库建立索引的知识图谱工具，能理解函数、类、路由、模块之间的调用与依赖关系。

**典型用途**：
- `search_graph`：按名称/语义搜索函数、类、变量；
- `trace_path`：追踪「谁调用某函数」「某函数调用谁」、数据流向、跨服务调用；
- `get_architecture`：查看项目模块划分、依赖、热点、架构边界；
- `query_graph`：用 Cypher 查询复杂多跳关系、内聚度、复杂度分析。

**适用场景**：架构梳理、调用链追踪、影响面分析、死代码识别。

---

## 3. confluence — Jira & Confluence 协同平台

**是什么**：对接 yunify 的 Jira 工单系统和 Confluence 知识库。

**典型用途**：
- **Jira**：查工单详情（`get_issue`）、搜索列表（`search`）、读字段/评论/附件（`download_attachments`）、查看工作流与阻塞关系；
- **Confluence**：读页面内容（`get_page`）、搜索（`search`）、查子页面/层级/标签/附件；
- 依赖环境里的 `JIRA_URL` + `JIRA_PERSONAL_TOKEN`、`CONFLUENCE_URL` + `CONFLUENCE_PERSONAL_TOKEN` 完成认证。

---

## 4. context7 — 库文档检索

**是什么**：检索最新、权威的第三方库 / 框架 / SDK 官方文档。

**典型用途**：
- 查 API 语法、配置方式、版本迁移、库特定调试、安装说明；
- 优先于凭记忆作答，避免训练数据过时；
- `resolve-library-id` 先解析出库 ID，再用 `query-docs` 查具体问题。

---

## 5. grep_app — GitHub 真实代码检索

**是什么**：从超过百万个公开 GitHub 仓库中检索真实代码示例（可理解为 GitHub 全局 grep）。

**典型用途**：
- 查某个库/API 的真实用法、生产级实现范例；
- 按语言、仓库、文件路径过滤；
- 支持正则匹配跨行模式，适合「别人怎么实现某功能」的问题。

---

## 6. lsp — 语言服务器协议集成

**是什么**：通过 LSP（Language Server Protocol）提供编辑器级的代码智能能力。

**典型用途**：
- 诊断：获取文件/目录的错误、警告、提示（`diagnostics`）；
- 跳转：找定义（`goto_definition`）、找引用（`find_references`）；
- 重命名符号、查文档/大纲符号；
- 用于改码后验证语法、做安全重构。

---

## 7. vision-ocr-mcp — 图片 / PDF / 剪贴板文字识别

**是什么**：基于 macOS Vision Framework 的 OCR 工具，从图片、PDF、屏幕截图、剪贴板中提取文字。

**典型用途**：
- `ocr_extract_text`（`file_path` / `base64` / `url`）解析图片或 PDF；
- `ocr_clipboard` 读取系统剪贴板中的图片文字；
- `ocr_screen_region` 截屏识别、`batch_ocr` 批量处理；
- `analyze_document` / `detect_text_regions` 获取文字位置与版面结构；
- `detect_barcodes` 识别二维码/条码；`watch_folder` 监控文件夹自动 OCR。

**注意**：AI 文本模型本身不具备视觉能力，遇到图片必须借助本工具解析，不能假装「看图」。

---

## 8. websearch — 网页搜索与抓取

**是什么**：联网搜索并按需抓取网页内容，转换为 Markdown 等格式供 AI 使用。

**典型用途**：
- 搜索最新信息、新闻、事实、人物、公司；
- 语义化描述理想页面来获得高质量结果（`websearch_web_search_exa`）；
- `webfetch` 抓取指定 URL 内容并转 Markdown。

---

## 9. wechat-devtools — 微信开发者工具（当前禁用）

**是什么**：对接微信小程序开发者工具的 MCP 服务。

**当前状态**：**Disabled（未启用）**。若需调试/构建微信小程序，需先在 opencode 配置中启用并认证后使用。

---

# OpenCode Skills 功能说明

Skill（技能）是预置的工作流/专业指令，为特定任务注入领域知识和执行规范。调用 skill 后，AI 会加载对应的详细指令并按步骤执行。优先级：**项目 > 用户 > opencode > 内置**。

## 一、用户自定义 Skills（最强优先级）

| Skill | 作用 |
| --- | --- |
| `mx-data` | 基于东方财富权威数据库的金融数据查询（行情、财务、关联关系）。 |
| `mx-search` | 东方财富妙想搜索，金融场景信源智能筛选（新闻、公告、研报、政策、事件）。 |
| `mx-xuangu` | 基于行情/财务指标选股，查询板块成分股、行业股票、个股/板块推荐，避免使用过时信息。 |
| `mx-zixuan` | 妙想自选管理：查询、添加、删除自选股（基于东方财富通行证账户）。 |
| `mx-moni` | 选股/相关命令入口。 |
| `grill-me` | 通过不间断追问，帮助打磨和完善方案或设计。 |
| `codebase-memory` | 用代码知识图谱做结构化代码查询（架构、调用链、影响分析）。 |

## 二、Superpowers 内置 Skills（工作流类）

| 分类 | Skill | 作用 |
| --- | --- | --- |
| 流程 | `using-superpowers` | 会话启动时规范如何使用 skills。 |
| 需求 | `brainstorming` | 任何创造性工作（建功能/组件/改行为）前先探索意图与需求。 |
| 计划 | `writing-plans` | 有规格需求时，写多步实施计划。 |
| 计划 | `executing-plans` | 分会话执行已写好的实施计划（含复核检查点）。 |
| 计划 | `dispatching-parallel-agents` | 面对 2+ 个无共享状态/无顺序依赖的独立任务时并行分发。 |
| 计划 | `subagent-driven-development` | 在当前会话用子代理执行实施计划中的独立任务。 |
| 开发 | `test-driven-development` | 实现功能/修 bug 前先写测试（TDD）。 |
| 开发 | `using-git-worktrees` | 特性开发前建立隔离的工作区（git worktree）。 |
| 调试 | `systematic-debugging` | 遇到 bug/测试失败/异常行为时，系统化假设-验证-定位根因。 |
| 完成 | `verification-before-completion` | 声称完成/修复前先运行验证命令并确认输出，证据先于断言。 |
| 完成 | `finishing-a-development-branch` | 实现完成、测试通过后，选择合并/提 PR/清理的收尾方式。 |
| 评审 | `requesting-code-review` | 完成任务/大功能/合并前，验证工作符合需求。 |
| 评审 | `receiving-code-review` | 收到评审反馈后，先验证质疑再进行修改。 |
| 安全 | `security-research` / `security-review` | 团队模式安全审计：并行猎人 + PoC 工程师验证可利用性、定级。 |
| 沉淀 | `writing-skills` | 创建/编辑/验证新技能。 |

## 三、常用内置命令（Slash Commands）

| 命令 | 作用 |
| --- | --- |
| `/mx-data` `/mx-search` `/mx-xuangu` `/mx-zixuan` | 调用对应金融类 skill。 |
| `/codebase-memory` | 调用代码知识图谱 skill。 |
| `/grill-me` | 用追问打磨方案/设计。 |
| `/debugging` | 任何语言的真实运行时调试（崩溃、静默失败、内存泄漏、卡死等）。 |
| `/refactor` | 智能重构（配合 LSP、AST-grep、TDD 验证）。 |
| `/git-master` | 任何 git 操作（原子提交、rebase/squash、历史检索 blame/bisect/log -S）。 |
| `/playwright` | 浏览器自动化任务（验证、信息抓取、截图、测试）。 |
| `/frontend-ui-ux` | 无设计稿时也能产出精美 UI/UX。 |
| `/visual-qa` | 构建/改完 UI 后做严谨的可视化 QA（网页 & 终端 UI）。 |
| `/review-work` | 并行 5 个子代理做目标/质量/安全/QA/上下文的全方位审查。 |
| `/security-research` `/security-review` | 团队模式安全审计。 |
| `/remove-ai-slops` | 移除 AI 生成代码的坏味道（AI slop）并按质量门槛验证。 |
| `/handoff` | 生成详细上下文摘要，便于在新会话继续工作。 |
| `/hyperplan` | 对抗式多代理规划（5 个对抗成员交叉批评，主导综合）。 |
| `/start-work` | 从 Prometheus 计划开始 Sisyphus 工作会话。 |
| `/init-deep` | 初始化分层 AGENTS.md 知识库。 |
| `/ralph-loop` `/ulw-loop` `/cancel-ralph` `/stop-continuation` | 自引用开发循环 / 超长工作循环 / 取消与停止。 |

---

---

# OpenCode 已安装插件

插件（Plugin）是 opencode 的功能扩展，提供额外的技能、命令、代理等工作流。在 `~/.config/opencode/opencode.jsonc` 的 `plugin` 字段中配置。

| # | 插件 | 来源 | 说明 |
| --- | --- | --- | --- |
| 1 | `superpowers` | GitHub 仓库 `jnMetaCode/superpowers-zh@v5.1.0` | Superpowers 系列工作流技能（上一章 Skills 中列出的 brainstorming、TDD、systematic-debugging 等均来自该插件）。 |
| 2 | `oh-my-openagent` | npm `@latest` | openagent 增强插件，提供额外的代理/命令能力。 |

**补充**：
- 以上两个插件的技能统一从本项目路径加载，例如 `superpowers/skills/`（详见 Skills 章节）。
- `package.json` 中还依赖 `oh-my-opencode`（版本 `^4.10.0`）与 `@opencode-ai/plugin`（`1.17.5`），服务于 opencode 的插件运行环境。
- 查看命令入口：`opencode` 会话内的 slash 命令（`/debugging`、`/refactor`、`/visual-qa` 等）多由上述插件注册提供。

## Skill 使用规则要点

1. **判断是否该用**：只要有一丁点可能适用，就必须调用 skill 检查，不能跳过。
2. **流程类优先**：先加载流程类 skill（如 brainstorming、debugging）决定「怎么做」，再加载实现类 skill 指导「怎么执行」。
3. **来源优先级**：项目 > 用户 > opencode > 内置。用户安装的 skill 覆盖内置默认。
4. **用户指令优先于 skill**：用户明确的指示（CLAUDE.md/AGENTS.md/直接请求）始终优先级最高。
