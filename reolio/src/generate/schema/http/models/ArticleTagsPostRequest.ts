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
 * @interface ArticleTagsPostRequest
 */
export interface ArticleTagsPostRequest {
  /**
   *
   * @type {string}
   * @memberof ArticleTagsPostRequest
   */
  name: string;
}

/**
 * Check if a given object implements the ArticleTagsPostRequest interface.
 */
export function instanceOfArticleTagsPostRequest(value: object): boolean {
  if (!("name" in value)) return false;
  return true;
}

export function ArticleTagsPostRequestFromJSON(
  json: any,
): ArticleTagsPostRequest {
  return ArticleTagsPostRequestFromJSONTyped(json, false);
}

export function ArticleTagsPostRequestFromJSONTyped(
  json: any,
  ignoreDiscriminator: boolean,
): ArticleTagsPostRequest {
  if (json == null) {
    return json;
  }
  return {
    name: json["name"],
  };
}

export function ArticleTagsPostRequestToJSON(
  value?: ArticleTagsPostRequest | null,
): any {
  if (value == null) {
    return value;
  }
  return {
    name: value["name"],
  };
}
