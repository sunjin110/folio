import { WordDetailWithTranslation } from "@/domain/model/word_detail_with_translation";
import { EnglishDictionaryRepository } from "@/domain/repository/english_dictionary";

export interface EnglishDictionaryUsecase {
  GetWordDetialWithTranslation(
    word: string,
  ): Promise<WordDetailWithTranslation>;
}

export function NewEnglishDictionaryUsecase(
  englishDictionaryRepo: EnglishDictionaryRepository,
): EnglishDictionaryUsecase {
  return new englishDictionary(englishDictionaryRepo);
}

class englishDictionary implements EnglishDictionaryUsecase {
  private englishDictionaryRepo: EnglishDictionaryRepository;

  constructor(englishDictionaryRepo: EnglishDictionaryRepository) {
    this.englishDictionaryRepo = englishDictionaryRepo;
  }

  GetWordDetialWithTranslation(
    word: string,
  ): Promise<WordDetailWithTranslation> {
    return this.englishDictionaryRepo.GetWordDetailWithTranslation(word);
  }
}
