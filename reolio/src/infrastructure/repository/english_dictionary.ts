import { WordDetailWithTranslation } from "@/domain/model/word_detail_with_translation";
import { EnglishDictionaryRepository } from "@/domain/repository/english_dictionary";
import { GolioApi } from "@/generate/schema/http";
import * as generate from "@/generate/schema/http";
import { wrapError } from "./error";
import { WordDetail } from "@/domain/model/word_detail";

export function NewEnglishDictionaryRepository(
  golioApi: GolioApi,
): EnglishDictionaryRepository {
  return new englishDictionary(golioApi);
}

class englishDictionary implements EnglishDictionaryRepository {
  private golioApi: GolioApi;
  constructor(golioApi: GolioApi) {
    this.golioApi = golioApi;
  }

  async GetWordDetailWithTranslation(
    word: string,
  ): Promise<WordDetailWithTranslation> {
    try {
      const resp = await this.golioApi.englishDictionaryWordGet({
        word: word,
      });
      return {
        origin: englishDictionary.toWordDetail(resp.origin),
        translated: englishDictionary.toWordDetail(resp.translated),
      };
    } catch (err) {
      throw wrapError(err, "failed get word from english dictionary");
    }
  }

  private static toWordDetail(wordDetail: generate.WordDetail): WordDetail {
    return {
      word: wordDetail.word,
      definitions: wordDetail.definitions,
      frequency: wordDetail.frequency,
    };
  }
}
