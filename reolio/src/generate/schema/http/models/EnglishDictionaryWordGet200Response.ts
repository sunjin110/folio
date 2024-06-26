/* tslint:disable */
/* eslint-disable */
/**
 * folio
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * The version of the OpenAPI document: 1.0.0
 *
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { mapValues } from "../runtime";
import type { WordDetail } from "./WordDetail";
import {
  WordDetailFromJSON,
  WordDetailFromJSONTyped,
  WordDetailToJSON,
} from "./WordDetail";

/**
 *
 * @export
 * @interface EnglishDictionaryWordGet200Response
 */
export interface EnglishDictionaryWordGet200Response {
  /**
   *
   * @type {WordDetail}
   * @memberof EnglishDictionaryWordGet200Response
   */
  origin: WordDetail;
  /**
   *
   * @type {WordDetail}
   * @memberof EnglishDictionaryWordGet200Response
   */
  translated: WordDetail;
}

/**
 * Check if a given object implements the EnglishDictionaryWordGet200Response interface.
 */
export function instanceOfEnglishDictionaryWordGet200Response(
  value: object,
): boolean {
  if (!("origin" in value)) return false;
  if (!("translated" in value)) return false;
  return true;
}

export function EnglishDictionaryWordGet200ResponseFromJSON(
  json: any,
): EnglishDictionaryWordGet200Response {
  return EnglishDictionaryWordGet200ResponseFromJSONTyped(json, false);
}

export function EnglishDictionaryWordGet200ResponseFromJSONTyped(
  json: any,
  ignoreDiscriminator: boolean,
): EnglishDictionaryWordGet200Response {
  if (json == null) {
    return json;
  }
  return {
    origin: WordDetailFromJSON(json["origin"]),
    translated: WordDetailFromJSON(json["translated"]),
  };
}

export function EnglishDictionaryWordGet200ResponseToJSON(
  value?: EnglishDictionaryWordGet200Response | null,
): any {
  if (value == null) {
    return value;
  }
  return {
    origin: WordDetailToJSON(value["origin"]),
    translated: WordDetailToJSON(value["translated"]),
  };
}
