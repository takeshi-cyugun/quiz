<script lang="ts">
  import { goto } from '$app/navigation';

  // 選択中のカテゴリー管理
  let selectedGenre = 'yoji';

  const genres = [
    {
      id: 'yoji',
      title: '四字熟語',
      description: '「意気投合」「以心伝心」など、漢字四文字の知恵をマスター', // Changed icon from empty to 📝
      icon: '👘'
    },
    {
      id: 'koto',
      title: 'ことわざ',
      description: '「猿も木から落ちる」など、先人の教えをクイズで学習',
      icon: '📜'
    }
  ];

  function handleNext() {
    goto(`/selection?genre=${selectedGenre}`);
  }
</script>

<main class="min-h-[100dvh] bg-slate-50 flex flex-col font-sans">
  <header class="bg-white border-b border-slate-200 px-4 py-6 text-center shrink-0">
    <h1 class="text-3xl font-black text-slate-900 tracking-tight">クイズアプリ</h1>
  </header>

  <div class="flex-1 overflow-y-auto p-6 flex flex-col items-center">
    <div class="w-full max-w-md">
      <div class="mb-8">
        <h2 class="text-2xl font-bold text-slate-900">ジャンルを選択</h2>
      </div>

      <div class="space-y-4">
        {#each genres as gen (gen.id)}
          <button
            on:click={() => selectedGenre = gen.id}
            class="w-full text-left p-6 rounded-2xl border-2 transition-all duration-200 flex items-center
              {selectedGenre === gen.id 
                ? 'border-indigo-600 bg-indigo-50 ring-4 ring-indigo-600/10' 
                : 'border-white bg-white shadow-sm hover:border-slate-200 active:scale-[0.98]'}"
          >
            <div class="text-4xl mr-5">{gen.icon}</div>
            <div class="flex-1">
              <h3 class="text-lg font-bold {selectedGenre === gen.id ? 'text-indigo-700' : 'text-slate-800'}">
                {gen.title}
              </h3>
              <p class="text-sm {selectedGenre === gen.id ? 'text-indigo-600/80' : 'text-slate-500'} leading-relaxed mt-1">
                {gen.description}
              </p>
            </div>
          </button>
        {/each}
      </div>

      <!-- 次へボタン -->
      <div class="mt-8 p-6 bg-white border-t border-slate-100 shrink-0 rounded-2xl shadow-sm">
        <button
          on:click={handleNext}
          class="w-full py-4 bg-indigo-600 text-white rounded-2xl font-bold text-lg shadow-lg shadow-indigo-200 hover:bg-indigo-700 active:scale-95 transition-all flex items-center justify-center space-x-2"
        >
          <span>クイズ選択へ進む</span>
          <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5" viewBox="0 0 20 20" fill="currentColor" aria-hidden="true">
            <path d="M7 4l6 6-6 6z" />
          </svg>
        </button>

      </div>
    </div>
  </div>
</main>