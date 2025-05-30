/* eslint-disable */
import * as types from './graphql';
import { TypedDocumentNode as DocumentNode } from '@graphql-typed-document-node/core';

/**
 * Map of all GraphQL operations in the project.
 *
 * This map has several performance disadvantages:
 * 1. It is not tree-shakeable, so it will include all operations in the project.
 * 2. It is not minifiable, so the string of a GraphQL query will be multiple times inside the bundle.
 * 3. It does not support dead code elimination, so it will add unused operations.
 *
 * Therefore it is highly recommended to use the babel or swc plugin for production.
 * Learn more about it here: https://the-guild.dev/graphql/codegen/plugins/presets/preset-client#reducing-bundle-size
 */
type Documents = {
    "\nmutation SendMessage($message: NewMessage!) {\n    sendMessage(message: $message) {\n        id\n        content\n        createdAt\n        owner\n    }\n}\n": typeof types.SendMessageDocument,
    "\nquery SearchMessages($query: String!, $channelID: String!) {\n    searchMessages(query:$query, channelID:$channelID) {\n        id\n        content\n        createdAt\n        owner\n        channelID\n    }\n}\n": typeof types.SearchMessagesDocument,
    "\nsubscription GetMessages($channelID: String!, $ownerID: String!) {\n    getMessages(channelID: $channelID, ownerID: $ownerID) {\n        id\n        content\n        createdAt\n        owner\n    }\n}\n": typeof types.GetMessagesDocument,
};
const documents: Documents = {
    "\nmutation SendMessage($message: NewMessage!) {\n    sendMessage(message: $message) {\n        id\n        content\n        createdAt\n        owner\n    }\n}\n": types.SendMessageDocument,
    "\nquery SearchMessages($query: String!, $channelID: String!) {\n    searchMessages(query:$query, channelID:$channelID) {\n        id\n        content\n        createdAt\n        owner\n        channelID\n    }\n}\n": types.SearchMessagesDocument,
    "\nsubscription GetMessages($channelID: String!, $ownerID: String!) {\n    getMessages(channelID: $channelID, ownerID: $ownerID) {\n        id\n        content\n        createdAt\n        owner\n    }\n}\n": types.GetMessagesDocument,
};

/**
 * The gql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 *
 *
 * @example
 * ```ts
 * const query = gql(`query GetUser($id: ID!) { user(id: $id) { name } }`);
 * ```
 *
 * The query argument is unknown!
 * Please regenerate the types.
 */
export function gql(source: string): unknown;

/**
 * The gql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function gql(source: "\nmutation SendMessage($message: NewMessage!) {\n    sendMessage(message: $message) {\n        id\n        content\n        createdAt\n        owner\n    }\n}\n"): (typeof documents)["\nmutation SendMessage($message: NewMessage!) {\n    sendMessage(message: $message) {\n        id\n        content\n        createdAt\n        owner\n    }\n}\n"];
/**
 * The gql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function gql(source: "\nquery SearchMessages($query: String!, $channelID: String!) {\n    searchMessages(query:$query, channelID:$channelID) {\n        id\n        content\n        createdAt\n        owner\n        channelID\n    }\n}\n"): (typeof documents)["\nquery SearchMessages($query: String!, $channelID: String!) {\n    searchMessages(query:$query, channelID:$channelID) {\n        id\n        content\n        createdAt\n        owner\n        channelID\n    }\n}\n"];
/**
 * The gql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function gql(source: "\nsubscription GetMessages($channelID: String!, $ownerID: String!) {\n    getMessages(channelID: $channelID, ownerID: $ownerID) {\n        id\n        content\n        createdAt\n        owner\n    }\n}\n"): (typeof documents)["\nsubscription GetMessages($channelID: String!, $ownerID: String!) {\n    getMessages(channelID: $channelID, ownerID: $ownerID) {\n        id\n        content\n        createdAt\n        owner\n    }\n}\n"];

export function gql(source: string) {
  return (documents as any)[source] ?? {};
}

export type DocumentType<TDocumentNode extends DocumentNode<any, any>> = TDocumentNode extends DocumentNode<  infer TType,  any>  ? TType  : never;
