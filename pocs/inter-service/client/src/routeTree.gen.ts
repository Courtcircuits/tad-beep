/* eslint-disable */

// @ts-nocheck

// noinspection JSUnusedGlobalSymbols

// This file was automatically generated by TanStack Router.
// You should NOT make any changes in this file as it will be overwritten.
// Additionally, you should also exclude this file from your linter and/or formatter to prevent it from being checked or modified.

// Import Routes

import { Route as rootRoute } from './routes/__root'
import { Route as IndexImport } from './routes/index'
import { Route as ChannelsIndexImport } from './routes/channels/index'
import { Route as ChannelsChannelIdImport } from './routes/channels/$channelId'

// Create/Update Routes

const IndexRoute = IndexImport.update({
  id: '/',
  path: '/',
  getParentRoute: () => rootRoute,
} as any)

const ChannelsIndexRoute = ChannelsIndexImport.update({
  id: '/channels/',
  path: '/channels/',
  getParentRoute: () => rootRoute,
} as any)

const ChannelsChannelIdRoute = ChannelsChannelIdImport.update({
  id: '/channels/$channelId',
  path: '/channels/$channelId',
  getParentRoute: () => rootRoute,
} as any)

// Populate the FileRoutesByPath interface

declare module '@tanstack/react-router' {
  interface FileRoutesByPath {
    '/': {
      id: '/'
      path: '/'
      fullPath: '/'
      preLoaderRoute: typeof IndexImport
      parentRoute: typeof rootRoute
    }
    '/channels/$channelId': {
      id: '/channels/$channelId'
      path: '/channels/$channelId'
      fullPath: '/channels/$channelId'
      preLoaderRoute: typeof ChannelsChannelIdImport
      parentRoute: typeof rootRoute
    }
    '/channels/': {
      id: '/channels/'
      path: '/channels'
      fullPath: '/channels'
      preLoaderRoute: typeof ChannelsIndexImport
      parentRoute: typeof rootRoute
    }
  }
}

// Create and export the route tree

export interface FileRoutesByFullPath {
  '/': typeof IndexRoute
  '/channels/$channelId': typeof ChannelsChannelIdRoute
  '/channels': typeof ChannelsIndexRoute
}

export interface FileRoutesByTo {
  '/': typeof IndexRoute
  '/channels/$channelId': typeof ChannelsChannelIdRoute
  '/channels': typeof ChannelsIndexRoute
}

export interface FileRoutesById {
  __root__: typeof rootRoute
  '/': typeof IndexRoute
  '/channels/$channelId': typeof ChannelsChannelIdRoute
  '/channels/': typeof ChannelsIndexRoute
}

export interface FileRouteTypes {
  fileRoutesByFullPath: FileRoutesByFullPath
  fullPaths: '/' | '/channels/$channelId' | '/channels'
  fileRoutesByTo: FileRoutesByTo
  to: '/' | '/channels/$channelId' | '/channels'
  id: '__root__' | '/' | '/channels/$channelId' | '/channels/'
  fileRoutesById: FileRoutesById
}

export interface RootRouteChildren {
  IndexRoute: typeof IndexRoute
  ChannelsChannelIdRoute: typeof ChannelsChannelIdRoute
  ChannelsIndexRoute: typeof ChannelsIndexRoute
}

const rootRouteChildren: RootRouteChildren = {
  IndexRoute: IndexRoute,
  ChannelsChannelIdRoute: ChannelsChannelIdRoute,
  ChannelsIndexRoute: ChannelsIndexRoute,
}

export const routeTree = rootRoute
  ._addFileChildren(rootRouteChildren)
  ._addFileTypes<FileRouteTypes>()

/* ROUTE_MANIFEST_START
{
  "routes": {
    "__root__": {
      "filePath": "__root.tsx",
      "children": [
        "/",
        "/channels/$channelId",
        "/channels/"
      ]
    },
    "/": {
      "filePath": "index.tsx"
    },
    "/channels/$channelId": {
      "filePath": "channels/$channelId.tsx"
    },
    "/channels/": {
      "filePath": "channels/index.tsx"
    }
  }
}
ROUTE_MANIFEST_END */
