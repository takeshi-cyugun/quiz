import { writable } from 'svelte/store';

export interface QuizResult {
  genre: string;
  mode: string;
  score: number;
  correct: number;
  total: number;
}

export const quizResult = writable<QuizResult>({
  genre: 'yoji',
  mode: 'meaning_4_choice',
  score: 0,
  correct: 0,
  total: 0
});