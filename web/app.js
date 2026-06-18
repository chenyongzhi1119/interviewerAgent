// ─── JD Presets ──────────────────────────────────────────────────────────────
const JD_PRESETS = {
  ai_compiler: `【岗位】AI 编译器工程师（社招）
【公司】某大型互联网 / AI 芯片公司

【岗位职责】
1. 负责 AI 编译器前端（Graph IR 设计）、中端（优化 Pass）及后端（Codegen / Schedule）全链路的设计与研发；
2. 基于 MLIR / LLVM 框架，针对主流模型（Transformer / LLM / VLM / Diffusion / MoE）在自研 AI 加速器上进行深度性能优化；
3. 研究并落地算子融合、内存分析、多面体调度（Polyhedral）、循环变换等高级编译优化技术；
4. 与芯片/RTL 团队联合制定 ISA、Tensor-Core 规格，将软件调度策略前置到硬件设计阶段；
5. 跟踪 MLIR / TVM / XLA / Triton 等主流编译器社区动态，推动先进技术在内部落地；
6. 编写高质量 C++ 代码，参与代码评审，建设编译器测试体系（正确性 + 性能 benchmark）。

【任职要求】
1. 本科及以上学历，计算机 / 微电子 / 人工智能相关专业，硕士优先；
2. 5 年以上编译器或高性能计算开发经验，其中至少 3 年直接参与 AI 编译器（TVM / MLIR / XLA / IREE 等）核心模块开发；
3. 精通 C++17/20，有扎实的模板元编程实战经验；能用 Python 快速实现原型 Pass；
4. 深入掌握 LLVM Pass pipeline、Loop/Vector/SLP 向量化、Memory Hierarchy 优化；
5. 熟悉 CUDA / PTX 编程，有 GPU 内核手写或性能调优经验者优先；
6. 了解深度学习模型结构（Attention / MoE / Diffusion 等）及量化、剪枝等轻量化技术；
7. 有开源编译器社区贡献经历（LLVM / MLIR / TVM upstream PR）优先。`,

  inference: `【岗位】LLM 推理加速工程师（社招）
【公司】某大型互联网公司 AI 平台部

【岗位职责】
1. 负责大语言模型（LLM）在 GPU / NPU 集群上的高效部署与推理性能优化；
2. 研究并落地推理加速技术：量化（INT8/INT4/FP8）、KV Cache 压缩、连续批处理（Continuous Batching）、投机采样（Speculative Decoding）等；
3. 对 vLLM / TensorRT-LLM / SGLang / lmdeploy 等主流推理框架进行二次开发与性能调优；
4. 编写高性能 CUDA Kernel，包括 FlashAttention、MLA、算子融合等；
5. 建立推理性能基准测试体系（latency / throughput / cost），定期与业界方案对比；
6. 与模型团队协作，针对新模型架构（MoE、Mamba、RWKV 等）设计专项推理优化方案；
7. 负责推理服务的分布式扩展（Tensor Parallelism / Pipeline Parallelism），保障生产环境高可用。

【任职要求】
1. 本科及以上学历，计算机 / 电子 / 数学等理工科专业；
2. 3 年以上 LLM 推理优化相关经验，熟悉 GPU 体系结构（Ampere / Hopper / Ada 架构）；
3. 精通 CUDA / PTX 编程，有 FlashAttention / PagedAttention / MLA 等算子手写或优化经验；
4. 熟悉 vLLM / TensorRT-LLM / Triton 至少一个框架的源码，有 upstream 贡献者优先；
5. 熟悉量化原理（AWQ / GPTQ / SmoothQuant / FP8 等），有生产级量化部署经验；
6. 良好的 Python / C++ 双语开发能力，能快速定位并解决复杂性能瓶颈；
7. 有千亿参数模型（如 GPT-4 级别）分布式推理部署经验者优先。`,

  ai_infra: `【岗位】AI Infra 工程师 · 大模型训练平台方向（社招）
【公司】某顶级互联网公司 AI 基础设施团队

【岗位职责】
1. 负责大规模分布式训练系统的设计与研发，支撑千亿 / 万亿参数级别模型的稳定高效训练；
2. 深入 PyTorch / JAX 底层，开发高性能训练算子和通信原语（NCCL / HCCL 优化）；
3. 研究并实现分布式并行策略：数据并行（DDP / FSDP）、张量并行、流水线并行、Expert 并行；
4. 开发训练稳定性保障工具：梯度监控、自动断点续训（Checkpointing）、异常检测与自愈；
5. 构建训练效率分析平台，定位并解决计算 / 通信 / 存储瓶颈（MFU 优化）；
6. 与模型算法团队深度协作，为预训练 / SFT / RLHF 等不同阶段提供专项基础设施支持；
7. 跟踪 Megatron-LM / DeepSpeed / FlashAttention / Liger Kernel 等社区最新进展并持续集成。

【任职要求】
1. 硕士及以上学历，计算机 / 电子信息 / 数学相关专业；
2. 5 年以上系统工程经验，其中 3 年以上大模型训练系统研发经验；
3. 精通 Python / C++ / CUDA，熟悉 PyTorch Autograd / Dispatcher / ATen 底层机制；
4. 深刻理解数据并行、张量并行、流水线并行原理，有千卡以上规模集群调优经验；
5. 熟悉 NCCL / MPI 通信库，有 All-Reduce / Ring-AllReduce / All-to-All 等通信优化经验；
6. 熟悉 GPU 硬件架构（SM / Warp / Shared Memory / NVLink / NVSwitch），能做 roofline 分析；
7. 有 Megatron-LM / DeepSpeed / FairScale 等框架核心贡献者经历优先。`,

  backend_go: `【岗位】后端开发工程师（Go 方向，P6-P7）
【公司】某大型互联网公司核心业务线

【岗位职责】
1. 负责高并发、大流量核心后端服务的架构设计、研发与性能优化（QPS 10w+ 级别）；
2. 构建稳定可靠的微服务体系，包括服务注册/发现、熔断、限流、链路追踪等基础设施；
3. 设计高性能数据存储方案，包括 MySQL 分库分表、Redis 集群、消息队列（Kafka/RocketMQ）；
4. 参与系统容量规划、性能基准建设，推动服务 SLA（P99 < 10ms）达标；
5. 制定团队技术规范，开展代码评审，推动工程质量和研发效率持续提升；
6. 负责服务的全链路可观测性建设（Metrics / Tracing / Logging）；
7. 与产品、算法、数据团队协作，推动 AI 能力在业务侧的落地。

【任职要求】
1. 本科及以上学历，计算机相关专业；
2. 5 年以上后端开发经验，3 年以上 Go 语言生产环境经验；
3. 深入理解 Go Runtime（goroutine 调度 / GC / 逃逸分析），能定位并解决 OOM / goroutine 泄漏等问题；
4. 熟悉分布式系统理论（CAP / BASE / Raft / 一致性哈希），有实际大规模分布式系统开发经验；
5. 熟悉 MySQL 索引原理、事务隔离级别、MVCC；熟悉 Redis 数据结构与持久化机制；
6. 熟悉 Kubernetes / Docker，有云原生服务部署和运维经验；
7. 良好的系统设计能力，能主导千万 DAU 级产品的技术方案设计与评审。`,

  pretrain: `【岗位】大模型预训练算法工程师（社招）
【公司】某大型互联网公司 AGI 团队

【岗位职责】
1. 负责大语言模型（LLM）/多模态大模型（VLM）的预训练核心算法研究与工程实现；
2. 研究 Scaling Laws，设计并验证不同规模模型的最优训练配方（数据配比、超参数、课程学习）；
3. 探索新型模型架构（MoE / MLA / 线性 Attention / SSM 等），推动训练效率与模型能力的帕累托提升；
4. 负责高质量预训练数据的获取、清洗、去重、质量评估 pipeline 建设（TB 级别 token 处理）；
5. 设计多阶段训练策略：预训练 → 持续预训练（CPT）→ 指令微调（SFT）→ 对齐（RLHF / DPO / GRPO）；
6. 建立模型能力评估体系（MMLU / MATH / HumanEval / 内部 benchmark），指导训练迭代方向；
7. 与 Infra 团队协作，在万卡 GPU 集群上稳定推进训练，MFU 持续优化。

【任职要求】
1. 硕士及以上学历，计算机 / 数学 / 统计 / 物理等相关专业，博士优先；
2. 3 年以上大模型预训练经验，独立主导过 70B 参数以上模型从零训练者优先；
3. 深入理解 Transformer 架构及其变体（GQA / RoPE / FlashAttention / SWA 等），能快速实现并验证新结构；
4. 熟悉主流训练框架（Megatron-LM / DeepSpeed），有 ZeRO / 张量并行 / 流水线并行调优经验；
5. 扎实的机器学习理论基础，能从 loss curve 诊断训练问题（spike / divergence / 过拟合）；
6. 有顶会论文发表（NeurIPS / ICML / ICLR / ACL）或参与过知名开源大模型项目者优先；
7. 强烈的研究驱动力和工程落地意识，能在快节奏环境中高效推进实验迭代。`,
};

function applyJDPreset(key) {
  if (!key) return;
  const textarea = document.getElementById('jd-input');
  textarea.value = JD_PRESETS[key] || '';
  if (window._validateStart) window._validateStart();
}

// ─── State ───────────────────────────────────────────────────────────────────
const state = {
  sessionId: null,
  company: null,
  provider: null,       // selected provider id
  round: 1,
  streaming: false,
  companies: [],
  providers: [],        // loaded from /api/providers
  // Rich file attachments (set during setup, carried into session)
  resumePDF: null,      // { data: base64, mime_type, name }
  resumeImages: [],     // [{ data: base64, mime_type, name }]
  jdImages: [],         // [{ data: base64, mime_type, name }]
  resumeMode: 'pdf',    // 'pdf' | 'image' | 'text'
  // Saved for next-round reuse
  savedJD: '',
  savedResume: '',
};

// ─── Init ─────────────────────────────────────────────────────────────────────
(async function init() {
  marked.setOptions({ breaks: true, gfm: true });
  await Promise.all([loadProviders(), loadCompanies()]);
  setupDragAndDrop();
  setupStartValidation();
})();

// Disable "开始面试" until provider + JD + resume are all present
function setupStartValidation() {
  const validate = () => {
    const hasProvider = !!state.provider;
    const hasJD = (document.getElementById('jd-input')?.value.trim().length > 0);
    const hasResume = state.resumeMode === 'pdf'
      ? document.getElementById('pdf-drop-zone')?.classList.contains('has-file')
      : (document.getElementById('resume-input')?.value.trim().length > 0);
    const btn = document.getElementById('start-btn');
    if (btn) btn.disabled = !(hasProvider && hasJD && hasResume);
  };
  // Watch relevant inputs
  document.getElementById('jd-input')?.addEventListener('input', validate);
  document.getElementById('resume-input')?.addEventListener('input', validate);
  // Also re-validate when provider changes (onProviderChange already fires)
  const origOnProviderChange = window.onProviderChange;
  // Validate on load
  setTimeout(validate, 100);
  // Export so PDF upload and provider change can re-trigger
  window._validateStart = validate;
}

async function loadCompanies() {
  try {
    const res = await fetch('/api/companies');
    const list = await res.json();
    state.companies = list || [];
    const sel = document.getElementById('company-select');
    sel.innerHTML = '';
    if (!list || list.length === 0) {
      sel.innerHTML = '<option value="">暂无公司模板</option>';
      return;
    }
    list.forEach(c => {
      const opt = document.createElement('option');
      opt.value = c.name;
      opt.textContent = c.display_name;
      sel.appendChild(opt);
    });
  } catch (e) {
    console.error('Failed to load companies', e);
  }
}

// ─── Settings (localStorage) ──────────────────────────────────────────────────
const SETTINGS_KEY = 'interviewer_provider_settings';

function loadStoredSettings() {
  try { return JSON.parse(localStorage.getItem(SETTINGS_KEY) || '{}'); }
  catch { return {}; }
}

function saveStoredSettings(obj) {
  localStorage.setItem(SETTINGS_KEY, JSON.stringify(obj));
}

function getProviderSetting(providerId) {
  return loadStoredSettings()[providerId] || {};
}

function openSettings() {
  const body = document.getElementById('settings-body');
  const stored = loadStoredSettings();
  body.innerHTML = '';

  state.providers.forEach(p => {
    const cfg = stored[p.id] || {};
    const hasKey = !!cfg.key;
    const statusClass = p.is_server_config ? 'server' : (hasKey ? 'local' : 'none');
    const statusText = p.is_server_config ? '服务器已配置' : (hasKey ? '已配置' : '未配置');

    const row = document.createElement('div');
    row.className = `provider-row${hasKey || p.is_server_config ? ' has-key' : ''}`;
    row.id = `prow-${p.id}`;
    row.innerHTML = `
      <div class="provider-row-header">
        <span class="provider-row-name">${p.display_name}</span>
        <span class="provider-status ${statusClass}" id="pstatus-${p.id}">${statusText}</span>
        ${p.register_url ? `<a class="provider-register" href="${p.register_url}" target="_blank">注册 →</a>` : ''}
      </div>
      <div class="key-row">
        <input
          class="key-input"
          id="pkey-${p.id}"
          type="password"
          placeholder="${p.is_server_config ? '（服务器已配置，可留空）' : 'API Key'}"
          value="${cfg.key || ''}"
          oninput="onKeyInput('${p.id}')"
          autocomplete="off"
        />
        <button class="key-clear" onclick="clearProviderKey('${p.id}')" title="清除">✕</button>
      </div>
      <span class="advanced-toggle" onclick="toggleAdvanced('${p.id}')">▸ 高级选项（自定义模型/地址）</span>
      <div class="advanced-fields" id="padv-${p.id}">
        <label>模型名称（留空使用默认：${p.model}）</label>
        <input id="pmodel-${p.id}" placeholder="${p.model}" value="${cfg.model || ''}" autocomplete="off" />
        ${p.default_base_url ? `
        <label>API 地址（留空使用默认）</label>
        <input id="pbase-${p.id}" placeholder="${p.default_base_url}" value="${cfg.base_url || ''}" autocomplete="off" />
        ` : ''}
      </div>
    `;
    body.appendChild(row);
  });

  document.getElementById('settings-modal').classList.add('open');
}

function toggleAdvanced(id) {
  const el = document.getElementById(`padv-${id}`);
  const toggle = el.previousElementSibling;
  const open = el.classList.toggle('open');
  toggle.textContent = (open ? '▾' : '▸') + ' 高级选项（自定义模型/地址）';
}

function onKeyInput(id) {
  const key = document.getElementById(`pkey-${id}`).value.trim();
  const row = document.getElementById(`prow-${id}`);
  const status = document.getElementById(`pstatus-${id}`);
  const prov = state.providers.find(p => p.id === id);
  if (key) {
    row.classList.add('has-key');
    status.className = 'provider-status local';
    status.textContent = '已配置';
  } else if (prov?.is_server_config) {
    row.classList.add('has-key');
    status.className = 'provider-status server';
    status.textContent = '服务器已配置';
  } else {
    row.classList.remove('has-key');
    status.className = 'provider-status none';
    status.textContent = '未配置';
  }
}

function clearProviderKey(id) {
  document.getElementById(`pkey-${id}`).value = '';
  onKeyInput(id);
}

function saveSettings() {
  const stored = {};
  state.providers.forEach(p => {
    const key = document.getElementById(`pkey-${p.id}`)?.value.trim() || '';
    const model = document.getElementById(`pmodel-${p.id}`)?.value.trim() || '';
    const base = document.getElementById(`pbase-${p.id}`)?.value.trim() || '';
    if (key || model || base) {
      stored[p.id] = { key, model, base_url: base };
    }
  });
  saveStoredSettings(stored);
  closeSettings();
  onProviderChange(); // refresh capability display
}

function closeSettings() {
  document.getElementById('settings-modal').classList.remove('open');
}

function handleModalClick(e) {
  if (e.target === document.getElementById('settings-modal')) closeSettings();
}

// ─── Providers ────────────────────────────────────────────────────────────────
async function loadProviders() {
  try {
    const res = await fetch('/api/providers');
    const list = await res.json();
    state.providers = list || [];
    const sel = document.getElementById('provider-select');
    sel.innerHTML = '';
    if (!list || list.length === 0) {
      sel.innerHTML = '<option value="">无可用供应商</option>';
      return;
    }
    list.forEach(p => {
      const opt = document.createElement('option');
      opt.value = p.id;
      opt.textContent = p.display_name;
      sel.appendChild(opt);
    });
    // Auto-select first
    state.provider = list[0].id;
    onProviderChange();
  } catch (e) {
    console.error('Failed to load providers', e);
  }
}

function onProviderChange() {
  const id = document.getElementById('provider-select').value;
  state.provider = id;
  const info = state.providers.find(p => p.id === id);
  if (!info) return;

  const cfg = getProviderSetting(id);
  const model = cfg.model || info.model;
  document.getElementById('model-display').textContent = model;

  // Show warning if provider is neither server-configured nor has a local key
  const hasKey = info.is_server_config || !!cfg.key;
  const startBtn = document.getElementById('start-btn');
  if (startBtn) {
    startBtn.title = hasKey ? '' : '请先在 ⚙️ 设置中填入此供应商的 API Key';
  }

  // PDF/image upload now works for all providers (extracted to text first)
  // No need to disable tabs based on provider capabilities
  if (window._validateStart) window._validateStart();
}

// Returns {key, model, base_url} to include in session creation request
function getProviderPayload() {
  const info = state.providers.find(p => p.id === state.provider);
  if (!info) return {};
  const cfg = getProviderSetting(state.provider);
  // Only send key if provider is NOT server-configured or user explicitly set one
  return {
    provider_key: cfg.key || '',
    provider_model: cfg.model || '',
    provider_base_url: cfg.base_url || '',
  };
}

// ─── Resume Mode Toggle ───────────────────────────────────────────────────────
function setResumeMode(mode) {
  state.resumeMode = mode;
  document.getElementById('resume-pdf-zone').style.display  = mode === 'pdf'  ? '' : 'none';
  // text zone needs flex+column to let textarea fill height
  document.getElementById('resume-text-zone').style.display = mode === 'text' ? 'flex' : 'none';
  document.getElementById('resume-text-zone').style.flexDirection = 'column';
  document.getElementById('resume-text-zone').style.flex = '1';
  document.getElementById('tab-pdf').classList.toggle('active',  mode === 'pdf');
  document.getElementById('tab-text').classList.toggle('active', mode === 'text');
}

// ─── PDF Upload ───────────────────────────────────────────────────────────────
function handlePDF(input) {
  const file = input.files[0];
  if (!file) return;
  const reader = new FileReader();
  reader.onload = async e => {
    const base64 = e.target.result.split(',')[1];
    setExtractStatus('pdf-label', `正在提取 ${file.name} 中的文字...`, 'loading');
    try {
      const text = await callExtract({ type: 'pdf', data: base64, mime_type: 'application/pdf' });
      // Switch to text mode and populate textarea
      setResumeMode('text');
      document.getElementById('resume-input').value = text;
      setExtractStatus('pdf-label', `✓ ${file.name} 已提取（可在文字模式中编辑）`, 'done');
      document.getElementById('pdf-drop-zone').classList.add('has-file');
      if (window._validateStart) window._validateStart();
    } catch (err) {
      setExtractStatus('pdf-label', `✗ 提取失败：${err.message}`, 'error');
    }
  };
  reader.readAsDataURL(file);
}

function clearPDF(e) {
  e.stopPropagation();
  state.resumePDF = null;
  document.getElementById('resume-file').value = '';
  document.getElementById('pdf-drop-zone').classList.remove('has-file');
  document.getElementById('pdf-label').textContent = '点击选择 PDF 简历，或拖拽到此处';
  document.getElementById('pdf-label').className = 'pdf-label';
}

function setupDragAndDrop() {
  const zone = document.getElementById('pdf-drop-zone');
  zone.addEventListener('dragover', e => { e.preventDefault(); zone.style.borderColor = 'var(--accent)'; });
  zone.addEventListener('dragleave', () => { zone.style.borderColor = ''; });
  zone.addEventListener('drop', e => {
    e.preventDefault();
    zone.style.borderColor = '';
    const file = e.dataTransfer.files[0];
    if (file && file.type === 'application/pdf') handlePDF({ files: [file] });
  });
}

function setExtractStatus(labelId, text, state) {
  const el = document.getElementById(labelId);
  if (!el) return;
  el.textContent = text;
  el.className = 'pdf-label' + (state === 'done' ? ' loaded' : state === 'error' ? ' error' : ' loading');
}

// ─── Resume Image Paste (now extracts to text) ───────────────────────────────
function handleResumePaste(e) {
  const items = e.clipboardData?.items;
  if (!items) return;
  for (const item of items) {
    if (item.type.startsWith('image/')) {
      e.preventDefault();
      const file = item.getAsFile();
      const reader = new FileReader();
      reader.onload = async ev => {
        const base64 = ev.target.result.split(',')[1];
        const textarea = document.getElementById('resume-input');
        const hint = document.getElementById('resume-extract-hint');
        if (hint) { hint.textContent = '准备识别...'; hint.style.display = ''; }
        try {
          const text = await callExtract({ type: 'image', data: base64, mime_type: item.type, _hintEl: hint });
          textarea.value += (textarea.value ? '\n\n' : '') + text;
          if (hint) { hint.textContent = '✓ 图片文字已提取（Tesseract 本地识别，无需 API）'; setTimeout(() => { hint.style.display = 'none'; }, 4000); }
        } catch (err) {
          if (hint) { hint.textContent = `✗ 识别失败：${err.message}`; }
        }
      };
      reader.readAsDataURL(file);
      return;
    }
  }
}

// ─── JD Image Paste (now extracts to text) ───────────────────────────────────
function handleJDPaste(e) {
  const items = e.clipboardData?.items;
  if (!items) return;
  for (const item of items) {
    if (item.type.startsWith('image/')) {
      e.preventDefault();
      const file = item.getAsFile();
      const reader = new FileReader();
      reader.onload = async ev => {
        const base64 = ev.target.result.split(',')[1];
        const textarea = document.getElementById('jd-input');
        const hint = document.getElementById('jd-extract-hint');
        if (hint) { hint.textContent = '准备识别...'; hint.style.display = ''; }
        try {
          const text = await callExtract({ type: 'image', data: base64, mime_type: item.type, _hintEl: hint });
          textarea.value += (textarea.value ? '\n\n' : '') + text;
          if (hint) { hint.textContent = '✓ 截图文字已提取（Tesseract 本地识别，无需 API）'; setTimeout(() => { hint.style.display = 'none'; }, 4000); }
        } catch (err) {
          if (hint) { hint.textContent = `✗ 识别失败：${err.message}`; }
        }
      };
      reader.readAsDataURL(file);
      return;
    }
  }
}

// ─── Extract Helper ───────────────────────────────────────────────────────────

// Extract text from an image using Tesseract.js.
// Uses Tesseract.recognize() directly — no persistent worker, avoids stale language data.
// Language data is fetched from Tesseract.js's default CDN (jsDelivr) and cached by the browser.
async function ocrImage(base64Data, mimeType, hintEl) {
  const dataURL = `data:${mimeType};base64,${base64Data}`;

  const result = await Tesseract.recognize(
    dataURL,
    'chi_sim+eng',    // simplified Chinese + English
    {
      logger: m => {
        if (!hintEl) return;
        if (m.status === 'loading tesseract core') {
          hintEl.textContent = '正在加载 Tesseract 核心...';
        } else if (m.status === 'loading language traineddata') {
          hintEl.textContent = '正在下载中文语言包（首次约 12MB，之后缓存）...';
        } else if (m.status === 'initializing api') {
          hintEl.textContent = '初始化识别引擎...';
        } else if (m.status === 'recognizing text') {
          const pct = Math.floor((m.progress || 0) * 100);
          hintEl.textContent = `识别中 ${pct}%...`;
        }
      },
    },
  );

  return result.data.text.trim();
}

// callExtract dispatches to:
//   PDF  → /api/extract (Go library, server-side)
//   Image → Tesseract.js (browser-side, no API key needed)
async function callExtract(params) {
  if (params.type === 'pdf') {
    const res = await fetch('/api/extract', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(params),
    });
    if (!res.ok) throw new Error(await res.text());
    const data = await res.json();
    return data.text;
  }

  if (params.type === 'image') {
    return await ocrImage(params.data, params.mime_type, params._hintEl ?? null);
  }

  throw new Error('unknown extract type: ' + params.type);
}

// ─── Setup → Chat ─────────────────────────────────────────────────────────────
async function startInterview() {
  const company = document.getElementById('company-select').value;
  const round = parseInt(document.getElementById('round-select').value);
  const jdText = document.getElementById('jd-input').value.trim();
  const resumeText = state.resumeMode === 'text'
    ? document.getElementById('resume-input').value.trim()
    : '';

  if (!state.provider) return alert('请选择 AI 供应商');
  if (!company) return alert('请选择目标公司');
  if (!jdText && state.jdImages.length === 0) return alert('请填写岗位 JD（文字或截图）');
  if (!resumeText && !state.resumePDF && state.resumeImages.length === 0)
    return alert('请提供简历（PDF、截图或文字）');

  const btn = document.getElementById('start-btn');
  btn.disabled = true;
  btn.textContent = '正在初始化...';

  // Save for next-round reuse
  state.savedJD = jdText;
  state.savedResume = resumeText;

  try {
    const body = {
      company,
      provider: state.provider,
      round,
      jd: jdText,
      resume: resumeText,
      resume_pdf: state.resumePDF || undefined,
      resume_images: state.resumeImages.length > 0 ? state.resumeImages : undefined,
      jd_images: state.jdImages.length > 0 ? state.jdImages : undefined,
      ...getProviderPayload(),
    };

    const res = await fetch('/api/sessions', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(body),
    });
    if (!res.ok) throw new Error(await res.text());
    const sess = await res.json();
    state.sessionId = sess.id;
    state.company = sess.company;
    state.round = sess.round;

    switchView('chat');
    // Deactivate top-nav buttons while in chat
    ['setup','history'].forEach(p => document.getElementById(`nav-${p}`)?.classList.remove('active'));
    updateChatHeader();
    document.getElementById('messages').innerHTML = '';

    await streamSSE(`/api/sessions/${state.sessionId}/start`, null, 'interviewer');
  } catch (e) {
    alert('启动失败：' + e.message);
  } finally {
    btn.disabled = false;
    btn.textContent = '开始面试 →';
  }
}

// ─── Chat ─────────────────────────────────────────────────────────────────────
async function sendMessage() {
  if (state.streaming) return;
  const input = document.getElementById('chat-input');
  const text = input.value.trim();
  if (!text) return;

  input.value = '';
  autoResize(input);

  appendMessage('candidate', text);
  await streamSSE(`/api/sessions/${state.sessionId}/chat`, { message: text }, 'interviewer');
}

function handleKey(e) {
  // Enter alone → send; Shift+Enter → newline
  if (e.key === 'Enter' && !e.shiftKey) {
    e.preventDefault();
    sendMessage();
  }
}

// ─── Voice Input ──────────────────────────────────────────────────────────────
let _recognition = null;
let _voiceActive = false;
let _voiceInterim = ''; // interim (not-yet-final) text shown while speaking

function initVoice() {
  const SpeechRecognition = window.SpeechRecognition || window.webkitSpeechRecognition;
  if (!SpeechRecognition) return null;

  const r = new SpeechRecognition();
  r.lang = 'zh-CN';          // primary language; browser auto-detects English words
  r.continuous = true;        // keep listening until stopped
  r.interimResults = true;    // show words as they're spoken

  r.onresult = (e) => {
    const textarea = document.getElementById('chat-input');
    let interim = '';
    let newFinal = '';
    for (let i = e.resultIndex; i < e.results.length; i++) {
      const t = e.results[i][0].transcript;
      if (e.results[i].isFinal) {
        newFinal += t;
      } else {
        interim += t;
      }
    }
    // Replace interim placeholder, then append final text
    const base = textarea.value.endsWith(_voiceInterim)
      ? textarea.value.slice(0, textarea.value.length - _voiceInterim.length)
      : textarea.value;
    _voiceInterim = interim;
    textarea.value = base + newFinal + interim;
    // If there was final text, clear interim buffer
    if (newFinal) _voiceInterim = interim;
    autoResize(textarea);
  };

  r.onerror = (e) => {
    if (e.error !== 'no-speech') console.warn('Speech error:', e.error);
    stopVoice();
  };

  r.onend = () => {
    // Auto-restart if user hasn't clicked stop
    if (_voiceActive) {
      try { r.start(); } catch (_) {}
    }
  };

  return r;
}

function toggleVoice() {
  if (_voiceActive) {
    stopVoice();
  } else {
    startVoice();
  }
}

function startVoice() {
  if (!_recognition) {
    _recognition = initVoice();
  }
  if (!_recognition) {
    alert('当前浏览器不支持语音输入，请使用 Chrome 或 Safari。');
    return;
  }
  _voiceActive = true;
  _voiceInterim = '';
  try { _recognition.start(); } catch (_) {}
  const btn = document.getElementById('voice-btn');
  if (btn) {
    btn.classList.add('recording');
    btn.title = '点击停止录音';
  }
}

function stopVoice() {
  _voiceActive = false;
  _voiceInterim = '';
  if (_recognition) {
    try { _recognition.stop(); } catch (_) {}
  }
  const btn = document.getElementById('voice-btn');
  if (btn) {
    btn.classList.remove('recording');
    btn.title = '语音输入';
  }
}

function autoResize(el) {
  el.style.height = 'auto';
  el.style.height = Math.min(el.scrollHeight, 160) + 'px';
}

// ─── Evaluate ─────────────────────────────────────────────────────────────────
async function endRound() {
  if (state.streaming) return;
  document.getElementById('end-round-btn').disabled = true;

  let evalText = '';
  const container = document.getElementById('evaluate-content');
  container.innerHTML = '<div class="typing-indicator"><span></span><span></span><span></span></div>';
  switchView('evaluate');
  ['setup','history'].forEach(p => document.getElementById(`nav-${p}`)?.classList.remove('active'));
  updateNextRoundBtn();

  try {
    await streamSSEToEl(
      `/api/sessions/${state.sessionId}/evaluate`,
      null,
      container,
      text => { evalText = text; },
    );
    container.innerHTML = marked.parse(evalText);
  } catch (e) {
    container.textContent = '评估失败：' + e.message;
  } finally {
    document.getElementById('end-round-btn').disabled = false;
  }
}

function updateNextRoundBtn() {
  const btn = document.getElementById('next-round-btn');
  if (state.round >= 3) {
    btn.textContent = '面试结束 · 重新开始';
    btn.onclick = backToSetup;
  } else {
    btn.textContent = `继续第 ${state.round + 1} 面 →`;
    btn.onclick = nextRound;
  }
}

async function nextRound() {
  if (state.round >= 3) { backToSetup(); return; }
  const newRound = state.round + 1;

  try {
    const body = {
      company: state.company,
      provider: state.provider,
      round: newRound,
      jd: state.savedJD,
      resume: state.savedResume,
      resume_pdf: state.resumePDF || undefined,
      resume_images: state.resumeImages.length > 0 ? state.resumeImages : undefined,
      jd_images: state.jdImages.length > 0 ? state.jdImages : undefined,
      ...getProviderPayload(),
    };
    const res = await fetch('/api/sessions', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(body),
    });
    if (!res.ok) throw new Error(await res.text());
    const sess = await res.json();
    state.sessionId = sess.id;
    state.round = newRound;

    switchView('chat');
    ['setup','history'].forEach(p => document.getElementById(`nav-${p}`)?.classList.remove('active'));
    updateChatHeader();
    document.getElementById('messages').innerHTML = '';

    await streamSSE(`/api/sessions/${state.sessionId}/start`, null, 'interviewer');
  } catch (e) {
    alert('切换失败：' + e.message);
  }
}

function backToSetup() {
  state.sessionId = null;
  navTo('setup');
}

// ─── Navigation ───────────────────────────────────────────────────────────────
function navTo(page) {
  const viewMap = { setup: 'setup', history: 'history', chat: 'chat', evaluate: 'evaluate', detail: 'history-detail', coding: 'coding' };
  const viewId = viewMap[page] ?? page;
  switchView(viewId);

  ['setup', 'history', 'coding'].forEach(p => {
    document.getElementById(`nav-${p}`)?.classList.toggle('active', p === page);
  });

  if (page === 'history') loadHistory();
  if (page === 'coding')  initCodingView();
}

// ─── History List ─────────────────────────────────────────────────────────────
async function loadHistory() {
  const container = document.getElementById('history-list');
  container.innerHTML = '<div class="history-empty">加载中...</div>';

  try {
    const res = await fetch('/api/sessions');
    const list = await res.json();

    // Filter: only show sessions that have at least one visible message
    const meaningful = (list || []).filter(s => s.message_count > 0);

    document.getElementById('history-count').textContent = `共 ${meaningful.length} 条记录`;

    if (meaningful.length === 0) {
      container.innerHTML = '<div class="history-empty">暂无历史记录，开始一场面试吧</div>';
      return;
    }

    const roundLabel = { 1: '一面', 2: '二面', 3: '三面' };
    const statusLabel = { chatting: '进行中', evaluating: '评估中', done: '已完成' };

    container.innerHTML = '';
    meaningful.forEach(sess => {
      const card = document.createElement('div');
      card.className = 'session-card';
      card.onclick = () => openHistoryDetail(sess.id);

      const companyInfo = state.companies.find(c => c.name === sess.company);
      const companyName = companyInfo?.display_name ?? sess.company;
      const round = roundLabel[sess.round] ?? `第${sess.round}面`;
      const status = statusLabel[sess.status] ?? sess.status;
      const statusClass = sess.status === 'done' ? 'status-done'
                        : sess.status === 'evaluating' ? 'status-evaluating'
                        : 'status-chatting';

      card.innerHTML = `
        <div class="session-card-header">
          <span class="session-round-badge">${round}</span>
          <span class="session-company">${companyName}</span>
          <div class="session-meta">
            <span class="session-status ${statusClass}">${status}</span>
            <span>${sess.message_count} 条消息</span>
            <span>${sess.created_at}</span>
          </div>
        </div>
        ${sess.preview ? `<div class="session-preview">${escapeHtml(sess.preview)}</div>` : ''}
      `;
      container.appendChild(card);
    });
  } catch (e) {
    container.innerHTML = `<div class="history-empty">加载失败：${e.message}</div>`;
  }
}

// ─── History Detail ───────────────────────────────────────────────────────────
async function openHistoryDetail(sessionId) {
  switchView('history-detail');
  document.getElementById('detail-messages').innerHTML = '<div class="history-empty">加载中...</div>';
  document.getElementById('detail-actions').innerHTML = '';

  try {
    const res = await fetch(`/api/sessions/${sessionId}`);
    if (!res.ok) throw new Error(await res.text());
    const sess = await res.json();

    const companyInfo = state.companies.find(c => c.name === sess.company);
    const companyName = companyInfo?.display_name ?? sess.company;
    const roundLabel = { 1: '一面', 2: '二面', 3: '三面' };
    const round = roundLabel[sess.round] ?? `第${sess.round}面`;

    document.getElementById('detail-title').textContent = `${companyName} · ${round}`;
    document.getElementById('detail-meta').textContent =
      new Date(sess.created_at).toLocaleString('zh-CN') + ' · ' + statusText(sess.status);

    // Show "继续面试" button for in-progress sessions
    const actions = document.getElementById('detail-actions');
    if (sess.status === 'chatting') {
      const btn = document.createElement('button');
      btn.className = 'btn btn-primary';
      btn.textContent = '继续面试 →';
      btn.onclick = () => resumeSession(sess);
      actions.appendChild(btn);
    }

    const container = document.getElementById('detail-messages');
    container.innerHTML = '';

    const visible = (sess.messages ?? []).filter(m => !m.is_hidden);
    if (visible.length === 0) {
      container.innerHTML = '<div class="history-empty">暂无对话内容</div>';
      return;
    }
    visible.forEach(m => {
      const role = m.role === 'assistant' ? 'interviewer' : 'candidate';
      appendMessageTo(container, role, m.content);
    });
    container.scrollTop = 0;
  } catch (e) {
    document.getElementById('detail-messages').innerHTML = `<div class="history-empty">加载失败：${e.message}</div>`;
  }
}

function statusText(status) {
  return { chatting: '进行中', evaluating: '评估中', done: '已完成' }[status] ?? status;
}

// resumeSession restores a persisted in-progress session into the chat view.
// The session's messages are replayed into the message list, and the input
// is re-enabled so the user can continue answering.
async function resumeSession(sess) {
  // Restore runtime state
  state.sessionId = sess.id;
  state.company   = sess.company;
  state.round     = sess.round;
  state.provider  = sess.provider;

  // Rebuild message bubbles from saved history
  const container = document.getElementById('messages');
  container.innerHTML = '';
  const visible = (sess.messages ?? []).filter(m => !m.is_hidden);
  visible.forEach(m => {
    const role = m.role === 'assistant' ? 'interviewer' : 'candidate';
    // Use appendMessageTo with the messages container
    appendMessageTo(container, role, m.content);
  });

  // Switch to chat view and restore UI state
  switchView('chat');
  ['setup', 'history'].forEach(p => document.getElementById(`nav-${p}`)?.classList.remove('active'));
  updateChatHeader();
  setSendDisabled(false);
  state.streaming = false;
  scrollToBottom();
}

function escapeHtml(s) {
  return s.replace(/&/g,'&amp;').replace(/</g,'&lt;').replace(/>/g,'&gt;');
}

function appendMessageTo(container, role, text) {
  const wrap = document.createElement('div');
  wrap.className = `msg-wrap ${role}`;
  const avatar = document.createElement('div');
  avatar.className = 'msg-avatar';
  avatar.textContent = role === 'interviewer' ? 'AI' : '我';
  const bubble = document.createElement('div');
  bubble.className = 'msg-bubble';
  bubble.innerHTML = marked.parse(text);
  wrap.appendChild(avatar);
  wrap.appendChild(bubble);
  container.appendChild(wrap);
}

// ─── SSE Streaming ───────────────────────────────────────────────────────────
async function streamSSE(url, body, role) {
  state.streaming = true;
  setSendDisabled(true);

  const bubble = appendMessage(role, '');
  const typingEl = appendTyping();
  let fullText = '';

  try {
    const fetchOpts = body
      ? { method: 'POST', headers: { 'Content-Type': 'application/json' }, body: JSON.stringify(body) }
      : { method: 'POST' };

    const res = await fetch(url, fetchOpts);
    if (!res.ok) throw new Error(await res.text());

    const reader = res.body.getReader();
    const decoder = new TextDecoder();
    let buf = '';

    while (true) {
      const { done, value } = await reader.read();
      if (done) break;
      buf += decoder.decode(value, { stream: true });
      const lines = buf.split('\n');
      buf = lines.pop();
      for (const line of lines) {
        if (!line.startsWith('data: ')) continue;
        const chunk = line.slice(6);
        if (chunk === '[DONE]') break;
        if (chunk.startsWith('[ERROR]')) {
          const msg = chunk.slice(7).trim();
          bubble.innerHTML = `<span style="color:var(--warning)">⚠️ ${escapeHtml(msg)}</span>`;
          scrollToBottom();
          break;
        }
        fullText += chunk;
        bubble.innerHTML = marked.parse(fullText);
        scrollToBottom();
      }
    }
  } catch (e) {
    bubble.textContent = '（请求失败：' + e.message + '）';
  } finally {
    typingEl.remove();
    state.streaming = false;
    setSendDisabled(false);
    scrollToBottom();
  }
}

async function streamSSEToEl(url, body, container, onUpdate) {
  const fetchOpts = body
    ? { method: 'POST', headers: { 'Content-Type': 'application/json' }, body: JSON.stringify(body) }
    : { method: 'POST' };

  const res = await fetch(url, fetchOpts);
  if (!res.ok) throw new Error(await res.text());

  const reader = res.body.getReader();
  const decoder = new TextDecoder();
  let buf = '';
  let fullText = '';

  while (true) {
    const { done, value } = await reader.read();
    if (done) break;
    buf += decoder.decode(value, { stream: true });
    const lines = buf.split('\n');
    buf = lines.pop();
    for (const line of lines) {
      if (!line.startsWith('data: ')) continue;
      const chunk = line.slice(6);
      if (chunk === '[DONE]') break;
      if (chunk.startsWith('[ERROR]')) break;
      fullText += chunk;
      container.innerHTML = marked.parse(fullText);
      onUpdate(fullText);
    }
  }
  onUpdate(fullText);
}

// ─── DOM Helpers ──────────────────────────────────────────────────────────────
function appendMessage(role, text) {
  const wrap = document.createElement('div');
  wrap.className = `msg-wrap ${role}`;

  const avatar = document.createElement('div');
  avatar.className = 'msg-avatar';
  avatar.textContent = role === 'interviewer' ? 'AI' : '我';

  const bubble = document.createElement('div');
  bubble.className = 'msg-bubble';
  if (text) bubble.innerHTML = marked.parse(text);

  wrap.appendChild(avatar);
  wrap.appendChild(bubble);
  document.getElementById('messages').appendChild(wrap);
  scrollToBottom();
  return bubble;
}

function appendTyping() {
  const wrap = document.createElement('div');
  wrap.className = 'msg-wrap interviewer';
  wrap.innerHTML = `
    <div class="msg-avatar">AI</div>
    <div class="msg-bubble">
      <div class="typing-indicator"><span></span><span></span><span></span></div>
    </div>`;
  document.getElementById('messages').appendChild(wrap);
  scrollToBottom();
  return wrap;
}

function scrollToBottom() {
  const el = document.getElementById('messages');
  el.scrollTop = el.scrollHeight;
}

function setSendDisabled(disabled) {
  document.getElementById('send-btn').disabled = disabled;
  document.getElementById('chat-input').disabled = disabled;
}

function switchView(name) {
  document.querySelectorAll('.view').forEach(v => v.classList.remove('active'));
  // Support both short names (setup→setup-view) and full IDs (history-detail→history-detail-view)
  const el = document.getElementById(`${name}-view`);
  if (el) el.classList.add('active');
}

function updateChatHeader() {
  const roundLabels = { 1: '一面', 2: '二面', 3: '三面' };
  const label = roundLabels[state.round] || `第${state.round}面`;
  document.getElementById('round-tag').textContent = label;
  const company = state.companies.find(c => c.name === state.company);
  document.getElementById('company-display').textContent = company ? company.display_name : state.company;
}


// ═══════════════════════════════════════════════════════════════
// CODING MODE
// ═══════════════════════════════════════════════════════════════

const codingState = {
  mode: 'algo',
  problem: null,
  currentLang: 'go',
  // AI Coding chat state
  aiMessages: [],
  aiStreaming: false,
  _inited: false,
};

// ─── Init ─────────────────────────────────────────────────────
function initCodingView() {
  if (codingState._inited) return;
  codingState._inited = true;
  // Always start in AI Coding mode — algo tab removed
  randomAIProblem();
  document.getElementById('ai-editor-wrap').style.display = 'flex';
}

// ─── Mode Switch ──────────────────────────────────────────────
function setCodingMode(mode) {
  codingState.mode = mode;
  document.getElementById('ctab-algo').classList.toggle('active', mode === 'algo');
  document.getElementById('ctab-ai').classList.toggle('active', mode === 'ai');
  document.getElementById('algo-toolbar').style.display  = mode === 'algo' ? 'flex' : 'none';
  document.getElementById('ai-toolbar').style.display    = mode === 'ai'   ? 'flex' : 'none';
  document.getElementById('algo-editor-wrap').style.display = mode === 'algo' ? 'flex' : 'none';
  document.getElementById('ai-editor-wrap').style.display   = mode === 'ai'   ? 'flex' : 'none';

  if (mode === 'algo') randomProblem();
  else                  randomAIProblem();
}

// ─── Problem Loading ──────────────────────────────────────────
function randomProblem() {
  const problems = typeof ALGO_PROBLEMS !== 'undefined' ? ALGO_PROBLEMS : [];
  if (!problems.length) return;
  const p = problems[Math.floor(Math.random() * problems.length)];
  loadAlgoProblem(p);
}

function randomAIProblem() {
  const problems = typeof AI_CODING_PROBLEMS !== 'undefined' ? AI_CODING_PROBLEMS : [];
  if (!problems.length) return;
  const p = problems[Math.floor(Math.random() * problems.length)];
  loadAIProblem(p);
}

function loadAlgoProblem(p) {
  codingState.problem = p;
  renderProblemPanel(p, 'algo');
  const lcLink = document.getElementById('lc-link');
  if (lcLink) {
    lcLink.href = p.leetcode_url || `https://leetcode.cn/search/?q=${encodeURIComponent(p.title)}`;
  }
}

function loadAIProblem(p) {
  codingState.problem = p;
  codingState.aiMessages = [];
  renderProblemPanel(p, 'ai');
  // Start the AI Coding conversation
  initAICodingChat(p);
}

function renderProblemPanel(p, mode) {
  const panel = document.getElementById('problem-panel');
  if (mode === 'algo') {
    const descHTML = marked.parse(p.description || '');
    panel.innerHTML = `
      <div class="prob-header">
        <span class="prob-title">${escapeHtml(p.title)}</span>
        <span class="prob-diff diff-${p.difficulty}">${p.difficulty}</span>
      </div>
      <div class="prob-tags">${(p.tags||[]).map(t=>`<span class="prob-tag">${escapeHtml(t)}</span>`).join('')}</div>
      <div class="prob-companies">常见于：${(p.companies||[]).join('、')}</div>
      <div class="prob-desc">${descHTML}</div>
      ${p.examples.length ? `<div>
        <div class="prob-section-title">示例</div>
        ${p.examples.map(e=>`<div class="prob-example">输入：${e.input}\n输出：${e.output}${e.explanation?'\n说明：'+e.explanation:''}</div>`).join('')}
      </div>` : ''}
      ${p.constraints.length ? `<div>
        <div class="prob-section-title">约束条件</div>
        <div class="prob-constraints">${p.constraints.map(c=>`<div class="prob-constraint">${escapeHtml(c)}</div>`).join('')}</div>
      </div>` : ''}
    `;
  } else {
    panel.innerHTML = `
      <div class="prob-header">
        <span class="prob-title">${escapeHtml(p.title)}</span>
        <span class="prob-diff diff-${p.difficulty}">${p.difficulty}</span>
      </div>
      <div class="prob-companies">场景：${escapeHtml(p.company)} · ${escapeHtml(p.scenario)}</div>
      <div class="prob-tags">${(p.tags||[]).map(t=>`<span class="prob-tag">${escapeHtml(t)}</span>`).join('')}</div>
      <div><div class="prob-section-title">业务背景</div><div class="ai-prob-background">${escapeHtml(p.background)}</div></div>
      <div><div class="prob-section-title">功能要求</div>
        <div class="ai-req-list">${(p.requirements||[]).map(r=>`<div class="ai-req-item">${escapeHtml(r)}</div>`).join('')}</div>
      </div>
      <div style="font-size:12px;color:var(--text3)">技术栈：${escapeHtml(p.tech_stack||'')}</div>
      <div><div class="prob-section-title">考察维度</div>
        <div class="ai-eval-list">${(p.evaluation_points||[]).map(e=>`<div class="ai-eval-item">${marked.parse(e)}</div>`).join('')}</div>
      </div>
    `;
  }
}


// ─── AI Coding: Conversation ──────────────────────────────────
async function initAICodingChat(p) {
  codingState.aiMessages = [];
  codingState.aiStreaming = false;

  const msgContainer = document.getElementById('coding-ai-messages');
  msgContainer.innerHTML = '';
  document.getElementById('coding-ai-placeholder') && (document.getElementById('coding-ai-placeholder').remove());

  // AI sends first message to start the interview
  const systemPrompt = buildAICodingSystem(p);
  const openMsg = `请开始这道 AI Coding 面试题，向候选人简要介绍题目背景和你的期望，然后提出第一个问题（不要一次性给出所有问题）。直接开始，不要说"好的"之类的开场白。`;
  codingState.aiMessages = [{ role: 'user', content: openMsg }];

  const bubble = appendAICodingMsg('interviewer', '');
  appendAICodingTyping();

  await callLLMStream(
    systemPrompt,
    codingState.aiMessages,
    () => { removeAICodingTyping(); return bubble; },
    chunk => { bubble.innerHTML = marked.parse(chunk); scrollAIChat(); }
  );

  // Save the AI's first reply
  const aiText = bubble.innerHTML ? bubble.textContent || '' : '';
  codingState.aiMessages.push({ role: 'assistant', content: bubble.textContent || bubble.innerText || '' });
}

async function sendAICodingMessage() {
  if (codingState.aiStreaming) return;
  const input = document.getElementById('coding-ai-input');
  const text = input.value.trim();
  if (!text) return;
  input.value = '';

  appendAICodingMsg('candidate', text);
  codingState.aiMessages.push({ role: 'user', content: text });

  const bubble = appendAICodingMsg('interviewer', '');
  const typing = appendAICodingTyping();
  codingState.aiStreaming = true;
  document.getElementById('coding-ai-send').disabled = true;
  document.getElementById('coding-ai-input').disabled = true;

  const systemPrompt = buildAICodingSystem(codingState.problem);
  await callLLMStream(
    systemPrompt,
    codingState.aiMessages,
    () => { typing.remove(); return bubble; },
    chunk => { bubble.innerHTML = marked.parse(chunk); scrollAIChat(); }
  );

  codingState.aiMessages.push({ role: 'assistant', content: bubble.textContent || '' });
  codingState.aiStreaming = false;
  document.getElementById('coding-ai-send').disabled = false;
  document.getElementById('coding-ai-input').disabled = false;
}

function codingAIKey(e) {
  if (e.key === 'Enter' && !e.shiftKey) { e.preventDefault(); sendAICodingMessage(); }
}

function codingAIResize(el) {
  el.style.height = 'auto';
  el.style.height = Math.min(el.scrollHeight, 200) + 'px';
}

function buildAICodingSystem(p) {
  return `你是一名资深技术面试官，正在进行 AI Coding 专项面试，考察候选人利用 AI 工具（Copilot / ChatGPT / Claude 等）完成工程任务的能力。

【题目】${p.title}（${p.company} · ${p.scenario}）
【业务背景】${p.background}
【功能要求】${(p.requirements||[]).join('；')}
【技术栈】${p.tech_stack}

面试风格：
- 先让候选人描述思路，再引导他们展示 Prompt 设计
- 对候选人的每个回答进行 2-3 层追问（为什么这样设计？有没有考虑 XXX 边界？）
- 当候选人展示 Prompt 或代码时，从「Prompt 清晰度」「架构合理性」「代码质量感知」角度点评
- 每次只问一个问题，等候选人回答

考察维度：
${(p.evaluation_points||[]).join('\n')}

请用中文进行面试，语气专业严谨。`;
}

// ─── Shared: callLLMStream ────────────────────────────────────
// Calls /api/llm/stream and streams the response.
// onBubble(fullText) → returns the DOM element to update
// onChunk(fullText) → called on each chunk (optional)
async function callLLMStream(system, messages, onBubble, onChunk) {
  const cfg = getProviderPayload();
  const body = {
    system,
    messages,
    provider: state.provider || '',
    provider_key: cfg.provider_key || '',
    provider_model: cfg.provider_model || '',
    provider_base_url: cfg.provider_base_url || '',
  };

  try {
    const res = await fetch('/api/llm/stream', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(body),
    });
    if (!res.ok) throw new Error(await res.text());

    let bubble = null;
    let fullText = '';
    const reader = res.body.getReader();
    const decoder = new TextDecoder();
    let buf = '';

    while (true) {
      const { done, value } = await reader.read();
      if (done) break;
      buf += decoder.decode(value, { stream: true });
      const lines = buf.split('\n');
      buf = lines.pop();
      for (const line of lines) {
        if (!line.startsWith('data: ')) continue;
        const chunk = line.slice(6);
        if (chunk === '[DONE]') break;
        if (chunk.startsWith('[ERROR]')) {
          if (!bubble && onBubble) bubble = onBubble('');
          if (bubble) bubble.innerHTML = `<span style="color:var(--danger)">${escapeHtml(chunk.slice(7))}</span>`;
          break;
        }
        fullText += chunk;
        if (!bubble && onBubble) bubble = onBubble(fullText);
        if (onChunk) onChunk(fullText);
        else if (bubble) bubble.innerHTML = marked.parse(fullText);
      }
    }
    return fullText;
  } catch (e) {
    if (onBubble) {
      const el = onBubble('');
      if (el) el.innerHTML = `<span style="color:var(--danger)">⚠️ ${escapeHtml(e.message)}</span>`;
    }
    return '';
  }
}

// ─── AI Coding Chat DOM helpers ───────────────────────────────
function appendAICodingMsg(role, text) {
  const container = document.getElementById('coding-ai-messages');
  const placeholder = document.getElementById('coding-ai-placeholder');
  if (placeholder) placeholder.remove();

  const wrap = document.createElement('div');
  wrap.className = `msg-wrap ${role === 'interviewer' ? 'interviewer' : 'candidate'}`;

  const avatar = document.createElement('div');
  avatar.className = 'msg-avatar';
  avatar.textContent = role === 'interviewer' ? 'AI' : '你';

  const bubble = document.createElement('div');
  bubble.className = 'msg-bubble';
  if (text) bubble.innerHTML = marked.parse(text);

  wrap.appendChild(avatar);
  wrap.appendChild(bubble);
  container.appendChild(wrap);
  scrollAIChat();
  return bubble;
}

function appendAICodingTyping() {
  const container = document.getElementById('coding-ai-messages');
  const wrap = document.createElement('div');
  wrap.className = 'msg-wrap interviewer';
  wrap.id = 'coding-ai-typing';
  wrap.innerHTML = `<div class="msg-avatar">AI</div>
    <div class="msg-bubble"><div class="typing-indicator"><span></span><span></span><span></span></div></div>`;
  container.appendChild(wrap);
  scrollAIChat();
  return wrap;
}

function removeAICodingTyping() {
  const el = document.getElementById('coding-ai-typing');
  if (el) el.remove();
}

function scrollAIChat() {
  const c = document.getElementById('coding-ai-messages');
  if (c) c.scrollTop = c.scrollHeight;
}

