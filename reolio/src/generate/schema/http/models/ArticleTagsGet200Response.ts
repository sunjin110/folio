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
import type { ArticleTag } from "./ArticleTag";
import {
  ArticleTagFromJSON,
  ArticleTagFromJSONTyped,
  ArticleTagToJSON,
} from "./ArticleTag";

/**
 *
 * @export
 * @interface ArticleTagsGet200Response
 */
export interface ArticleTagsGet200Response {
  /**
   *
   * @type {Array<ArticleTag>}
   * @memberof ArticleTagsGet200Response
   */
  tags: Array<ArticleTag>;
}

/**
 * Check if a given object implements the ArticleTagsGet200Response interface.
 */
export function instanceOfArticleTagsGet200Response(value: object): boolean {
  if (!("tags" in value)) return false;
  return true;
}

export function ArticleTagsGet200ResponseFromJSON(
  json: any,
): ArticleTagsGet200Response {
  return ArticleTagsGet200ResponseFromJSONTyped(json, false);
}

export function ArticleTagsGet200ResponseFromJSONTyped(
  json: any,
  ignoreDiscriminator: boolean,
): ArticleTagsGet200Response {
  if (json == null) {
    return json;
  }
  return {
    tags: (json["tags"] as Array<any>).map(ArticleTagFromJSON),
  };
}

export function ArticleTagsGet200ResponseToJSON(
  value?: ArticleTagsGet200Response | null,
): any {
  if (value == null) {
    return value;
  }
  return {
    tags: (value["tags"] as Array<any>).map(ArticleTagToJSON),
  };
}