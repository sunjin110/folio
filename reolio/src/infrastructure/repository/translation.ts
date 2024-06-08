import { TranslationRepository } from "@/domain/repository/translation";
import { AuthError, InternalError } from "@/error/error";
import { GolioApi, LanguageCode, ResponseError } from "@/generate/schema/http";

export function NewTranslationRepository(golioApi: GolioApi): TranslationRepository {
    return new translation(golioApi);
}

class translation implements TranslationRepository {
    private golioApi: GolioApi;

    constructor(golioApi: GolioApi) {
        this.golioApi = golioApi;
    }

    async TranslationText(text: string, sourceLanguageCodeStr: string, targetLanguageCodeStr: string): Promise<string> {
        const sourceLanguageCode = translation.toLanguageCode(sourceLanguageCodeStr);
        if (!sourceLanguageCode) {
            throw new Error(`bad request sourceLanguageCode: ${sourceLanguageCodeStr}`);
        }

        const targetLanguageCode = translation.toLanguageCode(targetLanguageCodeStr);
        if (!targetLanguageCode) {
            throw new Error(`bad request targetLanguageCode: ${targetLanguageCodeStr}`);
        }

        let translatedText = '';
        try {
            const output = await this.golioApi.translationPost({
                translationPostRequest: {
                    sourceLanguageCode: sourceLanguageCode,
                    targetLanguageCode: targetLanguageCode,
                    text: text,
                }
            })
            translatedText = output.translatedText;
        } catch (err) {
            if (err instanceof ResponseError) {
                if (err.response.status === 401 || err.response.status === 403) {
                    throw new AuthError("failed translate text", err);
                }
                throw new InternalError("failed translate text", err);
            }
            throw err;
        }
        return translatedText;
    }

    private static toLanguageCode(str: string): LanguageCode | undefined {
        const values = Object.values(LanguageCode);
        if (values.includes(str as any)) {
            return str as LanguageCode;
        }
        return undefined;
    }
}
