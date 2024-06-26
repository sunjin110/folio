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
 * @interface Article
 */
export interface Article {
  /**
   *
   * @type {string}
   * @memberof Article
   */
  id: string;
  /**
   *
   * @type {string}
   * @memberof Article
   */
  title: string;
  /**
   *
   * @type {string}
   * @memberof Article
   */
  createdAt: string;
  /**
   *
   * @type {Array<ArticleTag>}
   * @memberof Article
   */
  tags: Array<ArticleTag>;
}

/**
 * Check if a given object implements the Article interface.
 */
export function instanceOfArticle(value: object): boolean {
  if (!("id" in value)) return false;
  if (!("title" in value)) return false;
  if (!("createdAt" in value)) return false;
  if (!("tags" in value)) return false;
  return true;
}

export function ArticleFromJSON(json: any): Article {
  return ArticleFromJSONTyped(json, false);
}

export function ArticleFromJSONTyped(
  json: any,
  ignoreDiscriminator: boolean,
): Article {
  if (json == null) {
    return json;
  }
  return {
    id: json["id"],
    title: json["title"],
    createdAt: json["created_at"],
    tags: (json["tags"] as Array<any>).map(ArticleTagFromJSON),
  };
}

export function ArticleToJSON(value?: Article | null): any {
  if (value == null) {
    return value;
  }
  return {
    id: value["id"],
    title: value["title"],
    created_at: value["createdAt"],
    tags: (value["tags"] as Array<any>).map(ArticleTagToJSON),
  };
}
