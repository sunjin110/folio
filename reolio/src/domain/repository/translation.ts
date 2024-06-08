export interface TranslationRepository {
    TranslationText(text: string, sourceLanguageCode: string, targetLanguageCode: string): Promise<string>;
}
