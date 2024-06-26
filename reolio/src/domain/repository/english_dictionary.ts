import { WordDetailWithTranslation } from "../model/word_detail_with_translation";

export interface EnglishDictionaryRepository {
  GetWordDetailWithTranslation(
    word: string,
  ): Promise<WordDetailWithTranslation>;
}
