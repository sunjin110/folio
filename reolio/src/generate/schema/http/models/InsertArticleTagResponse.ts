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
/**
 *
 * @export
 * @interface InsertArticleTagResponse
 */
export interface InsertArticleTagResponse {
  /**
   *
   * @type {string}
   * @memberof InsertArticleTagResponse
   */
  id: string;
}

/**
 * Check if a given object implements the InsertArticleTagResponse interface.
 */
export function instanceOfInsertArticleTagResponse(value: object): boolean {
  if (!("id" in value)) return false;
  return true;
}

export function InsertArticleTagResponseFromJSON(
  json: any,
): InsertArticleTagResponse {
  return InsertArticleTagResponseFromJSONTyped(json, false);
}

export function InsertArticleTagResponseFromJSONTyped(
  json: any,
  ignoreDiscriminator: boolean,
): InsertArticleTagResponse {
  if (json == null) {
    return json;
  }
  return {
    id: json["id"],
  };
}

export function InsertArticleTagResponseToJSON(
  value?: InsertArticleTagResponse | null,
): any {
  if (value == null) {
    return value;
  }
  return {
    id: value["id"],
  };
}
