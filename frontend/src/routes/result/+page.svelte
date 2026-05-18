<script lang="ts">
  import { goto } from '$app/navigation';
  import { quizResult } from '$lib/stores/quizStore';

  // リアクティブにストアの値を購読
  $: ({ genre, mode, score, correct, total } = $quizResult);

  $: genreLabel = genre === 'yoji' ? '四字熟語' : 'ことわざ';

  // スコアに応じたメッセージ
  $: resultMessage = (() => {
    if (score === 100) return '素晴らしい！全問正解です！';
    if (score >= 80) return 'その調子！高得点です！';
    if (score >= 40) return 'お疲れ様でした！次も頑張りましょう。';
    return '復習して再挑戦してみましょう。';
  })();
</script>

<main class="min-h-[100dvh] bg-slate-50 flex flex-col font-sans">
  <header class="bg-white border-b border-slate-200 px-4 py-6 text-center shrink-0">
    <h1 class="text-3xl font-black text-slate-900 tracking-tight">クイズアプリ</h1>
  </header>

  <div class="flex-1 overflow-y-auto p-6 flex flex-col items-center">
    <div class="w-full max-w-md text-center mt-4">
      <div class="mb-4">
        <span class="inline-block px-3 py-1 bg-indigo-100 text-indigo-700 text-xs font-bold rounded-full uppercase tracking-wider mb-2">
          RESULT
        </span>
        <h2 class="text-2xl font-bold text-slate-900">{genreLabel} クイズ終了</h2>
      </div>

      <!-- スコア表示 -->
      <div class="bg-white p-10 rounded-3xl shadow-xl shadow-slate-200/50 mb-8 border border-slate-100">
        <div class="text-sm font-bold text-slate-400 mb-1 uppercase">Your Score</div>
        <div class="text-7xl font-black text-indigo-600 mb-4">{score}<span class="text-2xl ml-1 text-slate-400">pt</span></div>
        <div class="text-lg font-medium text-slate-700 mb-2">{resultMessage}</div>
        <div class="text-slate-500">
          正解数: <span class="font-bold text-slate-800">{correct}</span> / {total}
        </div>
      </div>

      <!-- アクションボタン -->
      <div class="space-y-4">
        <button
          on:click={() => goto(`/quiz?genre=${genre}&mode=${mode}`)}
          class="w-full py-4 bg-indigo-600 text-white rounded-2xl font-bold text-lg shadow-lg shadow-indigo-200 hover:bg-indigo-700 active:scale-95 transition-all flex items-center justify-center space-x-2"
        >
          <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5" viewBox="0 0 20 20" fill="currentColor">
            <path fill-rule="evenodd" d="M4 2a1 1 0 011 1v2.101a7.002 7.002 0 0111.601 2.566 1 1 0 11-1.885.666A5.002 5.002 0 005.999 7H9a1 1 0 010 2H4a1 1 0 01-1-1V3a1 1 0 011-1z" clip-rule="evenodd" />
          </svg>
          <span>もう一度挑戦する</span>
        </button>

        <button
          on:click={() => goto(`/selection?genre=${genre}`)}
          class="w-full py-4 bg-white text-slate-700 border-2 border-slate-200 rounded-2xl font-bold text-lg hover:bg-slate-50 active:scale-95 transition-all flex items-center justify-center space-x-2"
        >
          <span>別の形式で遊ぶ</span>
        </button>

        <button
          on:click={() => goto('/')}
          class="w-full py-4 text-slate-500 text-sm font-medium hover:text-indigo-600 transition-colors flex items-center justify-center space-x-1"
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
