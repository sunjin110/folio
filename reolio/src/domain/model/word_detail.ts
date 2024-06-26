export interface WordDetail {
  word: string;
  definitions: WordDefinition[];
  frequency: number;
}

export interface WordDefinition {
  definition: string;
  partOfSpeech: string;
  synonyms: string[];
  antonyms: string[];
  examples: string[];
}
