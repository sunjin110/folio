/* tslint:disable */
/* eslint-disable */
/**
 * My Project
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * The version of the OpenAPI document: 1.0.0
 * 
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { mapValues } from '../runtime';
/**
 * 
 * @export
 * @interface ArticlesArticleIdAiPutRequest
 */
export interface ArticlesArticleIdAiPutRequest {
    /**
     * AIに対する命令のメッセージです
     * @type {string}
     * @memberof ArticlesArticleIdAiPutRequest
     */
    message: string;
}

/**
 * Check if a given object implements the ArticlesArticleIdAiPutRequest interface.
 */
export function instanceOfArticlesArticleIdAiPutRequest(value: object): boolean {
    if (!('message' in value)) return false;
    return true;
}

export function ArticlesArticleIdAiPutRequestFromJSON(json: any): ArticlesArticleIdAiPutRequest {
    return ArticlesArticleIdAiPutRequestFromJSONTyped(json, false);
}

export function ArticlesArticleIdAiPutRequestFromJSONTyped(json: any, ignoreDiscriminator: boolean): ArticlesArticleIdAiPutRequest {
    if (json == null) {
        return json;
    }
    return {
        
        'message': json['message'],
    };
}

export function ArticlesArticleIdAiPutRequestToJSON(value?: ArticlesArticleIdAiPutRequest | null): any {
    if (value == null) {
        return value;
    }
    return {
        
        'message': value['message'],
    };
}
