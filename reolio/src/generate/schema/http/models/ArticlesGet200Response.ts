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
import type { Article } from "./Article";
import {
  ArticleFromJSON,
  ArticleFromJSONTyped,
  ArticleToJSON,
} from "./Article";

/**
 *
 * @export
 * @interface ArticlesGet200Response
 */
export interface ArticlesGet200Response {
  /**
   *
   * @type {Array<Article>}
   * @memberof ArticlesGet200Response
   */
  articles: Array<Article>;
  /**
   *
   * @type {number}
   * @memberof ArticlesGet200Response
   */
  total: number;
}

/**
 * Check if a given object implements the ArticlesGet200Response interface.
 */
export function instanceOfArticlesGet200Response(value: object): boolean {
  if (!("articles" in value)) return false;
  if (!("total" in value)) return false;
  return true;
}

export function ArticlesGet200ResponseFromJSON(
  json: any,
): ArticlesGet200Response {
  return ArticlesGet200ResponseFromJSONTyped(json, false);
}

export function ArticlesGet200ResponseFromJSONTyped(
  json: any,
  ignoreDiscriminator: boolean,
): ArticlesGet200Response {
  if (json == null) {
    return json;
  }
  return {
    articles: (json["articles"] as Array<any>).map(ArticleFromJSON),
    total: json["total"],
  };
}

export function ArticlesGet200ResponseToJSON(
  value?: ArticlesGet200Response | null,
): any {
  if (value == null) {
    return value;
  }
  return {
    articles: (value["articles"] as Array<any>).map(ArticleToJSON),
    total: value["total"],
  };
}
