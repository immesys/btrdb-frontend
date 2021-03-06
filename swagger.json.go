package main
const SwaggerJSON = `{
  "swagger": "2.0",
  "info": {
    "title": "BTrDB v4 API",
    "version": "4.10"
  },
  "schemes": [
    "http",
    "https"
  ],
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "paths": {
    "/v4/alignedwindows": {
      "post": {
        "summary": "Get aggregated power-of-two-length windows of data",
        "description": "AlignedWindows is similar to Windows but only permits a window length that is a power of two nanoseconds long, and starts on an exact multiple of the window length. It is faster for the database to execute, so should be used whenever the use case permits (e.g graphing)",
        "operationId": "AlignedWindows",
        "responses": {
          "200": {
            "description": "AlignedWindows has a streaming response, there will be multiple objects returned.",
            "schema": {
              "$ref": "#/definitions/grpcinterfaceAlignedWindowsResponse"
            },
            "examples":{
              "application/json":
              {
                "stat": {
                  "code": 0,
                  "msg": "dfdf",
                },
                "versionMajor": "451",
                "versionMinor": "0",
                "values": [
                  {
                    "time": "1495648275800604160",
                    "min": 10.1,
                    "mean": 11.5,
                    "max": 12.8,
                    "count": "242"
                  }
                ]
              }
            }

          }
        },
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/grpcinterfaceAlignedWindowsParams"
            }
          }
        ],
        "tags": [
          "BTrDB", "Read"
        ]
      }
    },
    "/v4/changes": {
      "post": {
        "summary": "Get the differences between two versions of a stream",
        "description": "Changes takes two versions and returns a sequence of time intervals where the streams have seen changes. For performance reasons, these windows can be 'rounded up', which is faster for the database to execute",
        "operationId": "Changes",
        "responses": {
          "200": {
            "description": "(streaming responses)",
            "schema": {
              "$ref": "#/definitions/grpcinterfaceChangesResponse"
            }
          }
        },
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/grpcinterfaceChangesParams"
            }
          }
        ],
        "tags": [
          "BTrDB"
        ]
      }
    },
    "/v4/create": {
      "post": {
        "operationId": "Create",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/grpcinterfaceCreateResponse"
            }
          }
        },
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/grpcinterfaceCreateParams"
            }
          }
        ],
        "tags": [
          "BTrDB"
        ]
      }
    },
    "/v4/delete": {
      "post": {
        "operationId": "Delete",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/grpcinterfaceDeleteResponse"
            }
          }
        },
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/grpcinterfaceDeleteParams"
            }
          }
        ],
        "tags": [
          "BTrDB"
        ]
      }
    },
    "/v4/faultinject": {
      "post": {
        "operationId": "FaultInject",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/grpcinterfaceFaultInjectResponse"
            }
          }
        },
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/grpcinterfaceFaultInjectParams"
            }
          }
        ],
        "tags": [
          "BTrDB"
        ]
      }
    },
    "/v4/flush": {
      "post": {
        "operationId": "Flush",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/grpcinterfaceFlushResponse"
            }
          }
        },
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/grpcinterfaceFlushParams"
            }
          }
        ],
        "tags": [
          "BTrDB"
        ]
      }
    },
    "/v4/info": {
      "post": {
        "operationId": "Info",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/grpcinterfaceInfoResponse"
            }
          }
        },
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/grpcinterfaceInfoParams"
            }
          }
        ],
        "tags": [
          "BTrDB"
        ]
      }
    },
    "/v4/insert": {
      "post": {
        "operationId": "Insert",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/grpcinterfaceInsertResponse"
            }
          }
        },
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/grpcinterfaceInsertParams"
            }
          }
        ],
        "tags": [
          "BTrDB"
        ]
      }
    },
    "/v4/listcollections": {
      "post": {
        "operationId": "ListCollections",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/grpcinterfaceListCollectionsResponse"
            }
          }
        },
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/grpcinterfaceListCollectionsParams"
            }
          }
        ],
        "tags": [
          "BTrDB"
        ]
      }
    },
    "/v4/lookupstreams": {
      "post": {
        "operationId": "LookupStreams",
        "responses": {
          "200": {
            "description": "(streaming responses)",
            "schema": {
              "$ref": "#/definitions/grpcinterfaceLookupStreamsResponse"
            }
          }
        },
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/grpcinterfaceLookupStreamsParams"
            }
          }
        ],
        "tags": [
          "BTrDB"
        ]
      }
    },
    "/v4/nearest": {
      "post": {
        "operationId": "Nearest",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/grpcinterfaceNearestResponse"
            }
          }
        },
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/grpcinterfaceNearestParams"
            }
          }
        ],
        "tags": [
          "BTrDB"
        ]
      }
    },
    "/v4/obliterate": {
      "post": {
        "operationId": "Obliterate",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/grpcinterfaceObliterateResponse"
            }
          }
        },
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/grpcinterfaceObliterateParams"
            }
          }
        ],
        "tags": [
          "BTrDB"
        ]
      }
    },
    "/v4/rawvalues": {
      "post": {
        "operationId": "RawValues",
        "responses": {
          "200": {
            "description": "(streaming responses)",
            "schema": {
              "$ref": "#/definitions/grpcinterfaceRawValuesResponse"
            }
          }
        },
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/grpcinterfaceRawValuesParams"
            }
          }
        ],
        "tags": [
          "BTrDB"
        ]
      }
    },
    "/v4/setstreamannotations": {
      "post": {
        "operationId": "SetStreamAnnotations",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/grpcinterfaceSetStreamAnnotationsResponse"
            }
          }
        },
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/grpcinterfaceSetStreamAnnotationsParams"
            }
          }
        ],
        "tags": [
          "BTrDB"
        ]
      }
    },
    "/v4/streaminfo": {
      "post": {
        "operationId": "StreamInfo",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/grpcinterfaceStreamInfoResponse"
            }
          }
        },
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/grpcinterfaceStreamInfoParams"
            }
          }
        ],
        "tags": [
          "BTrDB"
        ]
      }
    },
    "/v4/windows": {
      "post": {
        "operationId": "Windows",
        "responses": {
          "200": {
            "description": "(streaming responses)",
            "schema": {
              "$ref": "#/definitions/grpcinterfaceWindowsResponse"
            }
          }
        },
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/grpcinterfaceWindowsParams"
            }
          }
        ],
        "tags": [
          "BTrDB"
        ]
      }
    }
  },
  "definitions": {
    "grpcinterfaceAlignedWindowsParams": {
      "type": "object",
      "description":"top level desc",
      "properties": {
        "uuid": {
          "type": "string",
          "format": "byte",
          "description": "The stream UUID"
        },
        "start": {
          "type": "string",
          "format": "int64",
          "description": "The start time of the query in nanoseconds",
        },
        "end": {
          "type": "string",
          "format": "int64",
          "description": "The (exclusive) end time of the query in nanoseconds. This will be rounded down to the nearest window boundary"
        },
        "versionMajor": {
          "type": "string",
          "format": "uint64",
          "description": "The stream version to query"
        },
        "pointWidth": {
          "type": "integer",
          "format": "int64",
          "description": "The log (base 2) nanoseconds of the window width. E.g 34 equals 2^34 nanoseconds, roughly 17 seconds"
        }
      }
    },
    "grpcinterfaceAlignedWindowsResponse": {
      "type": "object",
      "properties": {
        "stat": {
          "$ref": "#/definitions/grpcinterfaceStatus"
        },
        "versionMajor": {
          "type": "string",
          "format": "uint64"
        },
        "versionMinor": {
          "type": "string",
          "format": "uint64"
        },
        "values": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/grpcinterfaceStatPoint"
          }
        }
      }
    },
    "grpcinterfaceChangedRange": {
      "type": "object",
      "properties": {
        "start": {
          "type": "string",
          "format": "int64"
        },
        "end": {
          "type": "string",
          "format": "int64"
        }
      }
    },
    "grpcinterfaceChangesParams": {
      "type": "object",
      "properties": {
        "uuid": {
          "type": "string",
          "format": "byte"
        },
        "fromMajor": {
          "type": "string",
          "format": "uint64"
        },
        "toMajor": {
          "type": "string",
          "format": "uint64"
        },
        "resolution": {
          "type": "integer",
          "format": "int64"
        }
      }
    },
    "grpcinterfaceChangesResponse": {
      "type": "object",
      "properties": {
        "stat": {
          "$ref": "#/definitions/grpcinterfaceStatus"
        },
        "versionMajor": {
          "type": "string",
          "format": "uint64"
        },
        "versionMinor": {
          "type": "string",
          "format": "uint64"
        },
        "ranges": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/grpcinterfaceChangedRange"
          }
        }
      }
    },
    "grpcinterfaceCreateParams": {
      "type": "object",
      "properties": {
        "uuid": {
          "type": "string",
          "format": "byte"
        },
        "collection": {
          "type": "string"
        },
        "tags": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/grpcinterfaceKeyValue"
          }
        },
        "annotations": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/grpcinterfaceKeyValue"
          }
        }
      }
    },
    "grpcinterfaceCreateResponse": {
      "type": "object",
      "properties": {
        "stat": {
          "$ref": "#/definitions/grpcinterfaceStatus"
        }
      }
    },
    "grpcinterfaceDeleteParams": {
      "type": "object",
      "properties": {
        "uuid": {
          "type": "string",
          "format": "byte"
        },
        "start": {
          "type": "string",
          "format": "int64"
        },
        "end": {
          "type": "string",
          "format": "int64"
        }
      }
    },
    "grpcinterfaceDeleteResponse": {
      "type": "object",
      "properties": {
        "stat": {
          "$ref": "#/definitions/grpcinterfaceStatus"
        },
        "versionMajor": {
          "type": "string",
          "format": "uint64"
        },
        "versionMinor": {
          "type": "string",
          "format": "uint64"
        }
      }
    },
    "grpcinterfaceFaultInjectParams": {
      "type": "object",
      "properties": {
        "type": {
          "type": "string",
          "format": "uint64"
        },
        "params": {
          "type": "string",
          "format": "byte"
        }
      }
    },
    "grpcinterfaceFaultInjectResponse": {
      "type": "object",
      "properties": {
        "stat": {
          "$ref": "#/definitions/grpcinterfaceStatus"
        },
        "rv": {
          "type": "string",
          "format": "byte"
        }
      }
    },
    "grpcinterfaceFlushParams": {
      "type": "object",
      "properties": {
        "uuid": {
          "type": "string",
          "format": "byte"
        }
      }
    },
    "grpcinterfaceFlushResponse": {
      "type": "object",
      "properties": {
        "stat": {
          "$ref": "#/definitions/grpcinterfaceStatus"
        }
      }
    },
    "grpcinterfaceInfoParams": {
      "type": "object"
    },
    "grpcinterfaceInfoResponse": {
      "type": "object",
      "properties": {
        "stat": {
          "$ref": "#/definitions/grpcinterfaceStatus"
        },
        "mash": {
          "$ref": "#/definitions/grpcinterfaceMash"
        },
        "majorVersion": {
          "type": "integer",
          "format": "int64"
        },
        "minorVersion": {
          "type": "integer",
          "format": "int64"
        },
        "build": {
          "type": "string"
        },
        "proxy": {
          "$ref": "#/definitions/grpcinterfaceProxyInfo"
        }
      }
    },
    "grpcinterfaceInsertParams": {
      "type": "object",
      "properties": {
        "uuid": {
          "type": "string",
          "format": "byte"
        },
        "sync": {
          "type": "boolean",
          "format": "boolean"
        },
        "values": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/grpcinterfaceRawPoint"
          }
        }
      }
    },
    "grpcinterfaceInsertResponse": {
      "type": "object",
      "properties": {
        "stat": {
          "$ref": "#/definitions/grpcinterfaceStatus"
        },
        "versionMajor": {
          "type": "string",
          "format": "uint64"
        },
        "versionMinor": {
          "type": "string",
          "format": "uint64"
        }
      }
    },
    "grpcinterfaceKeyOptValue": {
      "type": "object",
      "properties": {
        "key": {
          "type": "string"
        },
        "val": {
          "$ref": "#/definitions/grpcinterfaceOptValue"
        }
      }
    },
    "grpcinterfaceKeyValue": {
      "type": "object",
      "properties": {
        "key": {
          "type": "string"
        },
        "value": {
          "type": "string",
          "format": "byte"
        }
      }
    },
    "grpcinterfaceListCollectionsParams": {
      "type": "object",
      "properties": {
        "prefix": {
          "type": "string"
        },
        "startWith": {
          "type": "string"
        },
        "limit": {
          "type": "string",
          "format": "uint64"
        }
      }
    },
    "grpcinterfaceListCollectionsResponse": {
      "type": "object",
      "properties": {
        "stat": {
          "$ref": "#/definitions/grpcinterfaceStatus"
        },
        "collections": {
          "type": "array",
          "items": {
            "type": "string"
          }
        }
      }
    },
    "grpcinterfaceLookupStreamsParams": {
      "type": "object",
      "properties": {
        "collection": {
          "type": "string"
        },
        "isCollectionPrefix": {
          "type": "boolean",
          "format": "boolean"
        },
        "tags": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/grpcinterfaceKeyOptValue"
          }
        },
        "annotations": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/grpcinterfaceKeyOptValue"
          }
        }
      }
    },
    "grpcinterfaceLookupStreamsResponse": {
      "type": "object",
      "properties": {
        "stat": {
          "$ref": "#/definitions/grpcinterfaceStatus"
        },
        "results": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/grpcinterfaceStreamDescriptor"
          }
        }
      }
    },
    "grpcinterfaceMash": {
      "type": "object",
      "properties": {
        "revision": {
          "type": "string",
          "format": "int64"
        },
        "leader": {
          "type": "string"
        },
        "leaderRevision": {
          "type": "string",
          "format": "int64"
        },
        "totalWeight": {
          "type": "string",
          "format": "int64"
        },
        "healthy": {
          "type": "boolean",
          "format": "boolean"
        },
        "unmapped": {
          "type": "number",
          "format": "double"
        },
        "members": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/grpcinterfaceMember"
          }
        }
      }
    },
    "grpcinterfaceMember": {
      "type": "object",
      "properties": {
        "hash": {
          "type": "integer",
          "format": "int64"
        },
        "nodename": {
          "type": "string"
        },
        "up": {
          "type": "boolean",
          "format": "boolean"
        },
        "in": {
          "type": "boolean",
          "format": "boolean"
        },
        "enabled": {
          "type": "boolean",
          "format": "boolean"
        },
        "start": {
          "type": "string",
          "format": "int64"
        },
        "end": {
          "type": "string",
          "format": "int64"
        },
        "weight": {
          "type": "string",
          "format": "int64"
        },
        "readPreference": {
          "type": "number",
          "format": "double"
        },
        "httpEndpoints": {
          "type": "string"
        },
        "grpcEndpoints": {
          "type": "string"
        }
      }
    },
    "grpcinterfaceNearestParams": {
      "type": "object",
      "properties": {
        "uuid": {
          "type": "string",
          "format": "byte"
        },
        "time": {
          "type": "string",
          "format": "int64"
        },
        "versionMajor": {
          "type": "string",
          "format": "uint64"
        },
        "backward": {
          "type": "boolean",
          "format": "boolean"
        }
      }
    },
    "grpcinterfaceNearestResponse": {
      "type": "object",
      "properties": {
        "stat": {
          "$ref": "#/definitions/grpcinterfaceStatus"
        },
        "versionMajor": {
          "type": "string",
          "format": "uint64"
        },
        "versionMinor": {
          "type": "string",
          "format": "uint64"
        },
        "value": {
          "$ref": "#/definitions/grpcinterfaceRawPoint"
        }
      }
    },
    "grpcinterfaceObliterateParams": {
      "type": "object",
      "properties": {
        "uuid": {
          "type": "string",
          "format": "byte"
        }
      }
    },
    "grpcinterfaceObliterateResponse": {
      "type": "object",
      "properties": {
        "stat": {
          "$ref": "#/definitions/grpcinterfaceStatus"
        }
      }
    },
    "grpcinterfaceOptValue": {
      "type": "object",
      "properties": {
        "value": {
          "type": "string",
          "format": "byte"
        }
      }
    },
    "grpcinterfaceProxyInfo": {
      "type": "object",
      "properties": {
        "proxyEndpoints": {
          "type": "array",
          "items": {
            "type": "string"
          }
        }
      }
    },
    "grpcinterfaceRawPoint": {
      "type": "object",
      "properties": {
        "time": {
          "type": "string",
          "format": "int64"
        },
        "value": {
          "type": "number",
          "format": "double"
        }
      }
    },
    "grpcinterfaceRawValuesParams": {
      "type": "object",
      "properties": {
        "uuid": {
          "type": "string",
          "format": "byte"
        },
        "start": {
          "type": "string",
          "format": "int64"
        },
        "end": {
          "type": "string",
          "format": "int64"
        },
        "versionMajor": {
          "type": "string",
          "format": "uint64"
        }
      }
    },
    "grpcinterfaceRawValuesResponse": {
      "type": "object",
      "properties": {
        "stat": {
          "$ref": "#/definitions/grpcinterfaceStatus"
        },
        "versionMajor": {
          "type": "string",
          "format": "uint64"
        },
        "versionMinor": {
          "type": "string",
          "format": "uint64"
        },
        "values": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/grpcinterfaceRawPoint"
          }
        }
      }
    },
    "grpcinterfaceSetStreamAnnotationsParams": {
      "type": "object",
      "properties": {
        "uuid": {
          "type": "string",
          "format": "byte"
        },
        "expectedAnnotationVersion": {
          "type": "string",
          "format": "uint64"
        },
        "annotations": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/grpcinterfaceKeyOptValue"
          }
        }
      }
    },
    "grpcinterfaceSetStreamAnnotationsResponse": {
      "type": "object",
      "properties": {
        "stat": {
          "$ref": "#/definitions/grpcinterfaceStatus"
        }
      }
    },
    "grpcinterfaceStatPoint": {
      "type": "object",
      "properties": {
        "time": {
          "type": "string",
          "format": "int64"
        },
        "min": {
          "type": "number",
          "format": "double"
        },
        "mean": {
          "type": "number",
          "format": "double"
        },
        "max": {
          "type": "number",
          "format": "double"
        },
        "count": {
          "type": "string",
          "format": "uint64"
        }
      }
    },
    "grpcinterfaceStatus": {
      "type": "object",
      "properties": {
        "code": {
          "type": "integer",
          "format": "int64"
        },
        "msg": {
          "type": "string"
        },
        "mash": {
          "$ref": "#/definitions/grpcinterfaceMash"
        }
      }
    },
    "grpcinterfaceStreamDescriptor": {
      "type": "object",
      "properties": {
        "uuid": {
          "type": "string",
          "format": "byte"
        },
        "collection": {
          "type": "string"
        },
        "tags": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/grpcinterfaceKeyValue"
          }
        },
        "annotations": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/grpcinterfaceKeyValue"
          }
        },
        "annotationVersion": {
          "type": "string",
          "format": "uint64"
        }
      }
    },
    "grpcinterfaceStreamInfoParams": {
      "type": "object",
      "properties": {
        "uuid": {
          "type": "string",
          "format": "byte"
        },
        "omitVersion": {
          "type": "boolean",
          "format": "boolean"
        },
        "omitDescriptor": {
          "type": "boolean",
          "format": "boolean"
        }
      }
    },
    "grpcinterfaceStreamInfoResponse": {
      "type": "object",
      "properties": {
        "stat": {
          "$ref": "#/definitions/grpcinterfaceStatus"
        },
        "versionMajor": {
          "type": "string",
          "format": "uint64"
        },
        "versionMinor": {
          "type": "string",
          "format": "uint64"
        },
        "descriptor": {
          "$ref": "#/definitions/grpcinterfaceStreamDescriptor"
        }
      }
    },
    "grpcinterfaceWindowsParams": {
      "type": "object",
      "properties": {
        "uuid": {
          "type": "string",
          "format": "byte"
        },
        "start": {
          "type": "string",
          "format": "int64"
        },
        "end": {
          "type": "string",
          "format": "int64"
        },
        "versionMajor": {
          "type": "string",
          "format": "uint64"
        },
        "width": {
          "type": "string",
          "format": "uint64"
        },
        "depth": {
          "type": "integer",
          "format": "int64"
        }
      }
    },
    "grpcinterfaceWindowsResponse": {
      "type": "object",
      "properties": {
        "stat": {
          "$ref": "#/definitions/grpcinterfaceStatus"
        },
        "versionMajor": {
          "type": "string",
          "format": "uint64"
        },
        "versionMinor": {
          "type": "string",
          "format": "uint64"
        },
        "values": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/grpcinterfaceStatPoint"
          }
        }
      }
    }
  }
}
`;