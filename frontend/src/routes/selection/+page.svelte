<script lang="ts">
  import { page } from '$app/stores';
  import { goto } from '$app/navigation';

  // クエリパラメータからジャンルを取得 (yoji or koto)
  const genre = $page.url.searchParams.get('genre') || 'yoji';
  const genreLabel = genre === 'yoji' ? '四字熟語' : 'ことわざ';

  // 選択中のモード管理
  let selectedMode = 'meaning_4_choice';

  const modes = [
    {
      id: 'meaning_4_choice',
      title: '意味当て（4択）',
      description: `${genreLabel} の意味を4つの選択肢から正解を選びます。`,
      type: 'Lv1',
      icon: '📖'
    },
    {
      id: 'name_4_choice',
      title: '名称当て（4択）',
      description: `意味を読んで4つの ${genreLabel} から正解を選びます。`,
      type: 'Lv2',
      icon: '🎯'
    },
    {
      id: 'name_exact',
      title: '名称当て（記述）',
      description: `意味を読んで正解の ${genreLabel} を入力します。`,
      type: 'Lv3',
      icon: '✍️'
    }
  ];

  async function handleStart() {
    // 現時点ではモック動作としてコンソール出力し、クイズ画面（未作成）へ遷移
    console.log(`Starting quiz: Genre=${genre}, Mode=${selectedMode}`);
    goto(`/quiz?genre=${genre}&mode=${selectedMode}`);
  }
</script>

<main class="min-h-[100dvh] bg-slate-50 flex flex-col font-sans">
  <!-- ヘッダー兼ナビゲーション -->
  <header class="bg-white border-b border-slate-200 px-4 py-6 text-center shrink-0">
    <h1 class="text-3xl font-black text-slate-900 tracking-tight">クイズアプリ</h1>
  </header>

  <div class="flex-1 overflow-y-auto p-6 flex flex-col items-center">
    <div class="w-full max-w-md">
      <!-- カテゴリー表示 -->
      <div class="mb-8">
        <h2 class="text-2xl font-bold text-slate-900">{genreLabel} クイズ形式を選択</h2>
      </div>

      <!-- モードリスト -->
      <div class="space-y-4">
        {#each modes as mode (mode.id)}
          <button
            on:click={() => selectedMode = mode.id}
            class="w-full text-left p-6 rounded-2xl border-2 transition-all duration-200 flex items-center
              {selectedMode === mode.id 
                ? 'border-indigo-600 bg-indigo-50 ring-4 ring-indigo-600/10' 
                : 'border-white bg-white shadow-sm hover:border-slate-200 active:scale-[0.98]'}"
          >
            <div class="text-3xl mr-5">{mode.icon}</div>
            <div class="flex-1">
              <div class="flex items-center justify-between mb-1">
                <span class="text-lg font-bold {selectedMode === mode.id ? 'text-indigo-700' : 'text-slate-800'}">
                  {mode.title}
                </span>
                <span class="text-[10px] px-2 py-0.5 rounded font-bold border
                  {mode.type === 'Lv1' ? 'bg-green-100 text-green-800 border-green-200' : ''}
                  {mode.type === 'Lv2' ? 'bg-yellow-100 text-yellow-800 border-yellow-200' : ''}
                  {mode.type === 'Lv3' ? 'bg-red-100 text-red-800 border-red-200' : ''}
                ">
                  {mode.type}
                </span>
              </div>
              <p class="text-sm {selectedMode === mode.id ? 'text-indigo-600/80' : 'text-slate-500'} leading-relaxed">
                {mode.description}
              </p>
            </div>
          </button>
        {/each}
      </div>

      <!-- STARTボタン固定エリア -->
      <div class="mt-8 p-6 bg-white border-t border-slate-100 shrink-0 rounded-2xl shadow-sm">
        <button
          on:click={handleStart}
          class="w-full py-4 bg-indigo-600 text-white rounded-2xl font-bold text-lg shadow-lg shadow-indigo-200 hover:bg-indigo-700 active:scale-95 transition-all flex items-center justify-center space-x-2"
        >
          <span>クイズ開始</span>
          <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5" viewBox="0 0 20 20" fill="currentColor" aria-hidden="true">
            <path d="M7 4l6 6-6 6z" />
          </svg>
        </button>

        <button 
          on:click={() => goto('/')}
          class="w-full mt-4 text-slate-500 text-sm font-medium hover:text-indigo-600 transition-colors flex items-center justify-center space-x-1"
        >
          <svg xmlns="http://www.w3.org/2000/svg" class="h-4 w-4" viewBox="0 0 20 20" fill="currentColor" aria-hidden="true">
            <path d="M13 4l-6 6 6 6z" />
          </svg>
          <span>ジャンル選択に戻る</span>
        </button>
      </div>
    </div>
  </div>
</main>
