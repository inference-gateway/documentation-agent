# GetLibraryIDTool Implementation

This document describes the implementation of the GetLibraryIDTool (which is implemented as the `resolve_library_id` tool) for the Context7 A2A Agent.

## Overview

The `resolve_library_id` tool has been implemented to directly integrate with the Context7 API, providing library search and resolution capabilities similar to the official Context7 MCP server.

## Implementation Details

### Tools Implemented

#### 1. resolve_library_id
- **Purpose**: Resolves library names to Context7-compatible library IDs
- **Input Parameter**: `libraryName` (string) - Library name to search for
- **Output**: JSON response containing the selected library ID and metadata

#### 2. get_library_docs (Enhanced)
- **Purpose**: Fetches documentation for a specific library
- **Input Parameters**: 
  - `context7CompatibleLibraryID` (string, required) - Library ID in format `/org/project`
  - `tokens` (number, optional) - Maximum tokens to retrieve (default: 10000)
  - `topic` (string, optional) - Topic to focus documentation on
- **Output**: JSON response containing the documentation content

### Architecture

#### Direct API Integration
Instead of using MCP protocol communication, this implementation directly calls the Context7 REST APIs:

- **Search API**: `GET https://context7.com/api/v1/search?query={query}`
- **Documentation API**: `GET https://context7.com/api/v1/{libraryId}?tokens={tokens}&topic={topic}&type=txt`

#### Authentication
- Uses `Authorization: Bearer {apiKey}` header
- API key retrieved from `CONTEXT7_API_KEY` environment variable
- User-Agent set to `documentation-agent/0.1.0`

### Library Selection Logic

The `resolve_library_id` tool implements intelligent library selection based on the Context7 MCP server patterns:

1. **Exact Match Priority**: First looks for exact title matches (case-insensitive)
2. **Scoring System**: If no exact match, scores libraries based on:
   - Documentation coverage (`totalSnippets` count)
   - Trust score (7-10 range, weighted heavily)
   - State (prioritizes "finalized" libraries)
3. **Fallback**: Selects the first result if no scoring applies

### Response Format

#### resolve_library_id Response
```json
{
  "selectedLibraryID": "/vercel/next.js",
  "selectedLibrary": {
    "id": "/vercel/next.js",
    "title": "Next.js",
    "description": "The React Framework for Production",
    "totalSnippets": 150,
    "totalTokens": 50000,
    "state": "finalized",
    "lastUpdateDate": "2025-01-15T10:30:00Z",
    "trustScore": 9,
    "stars": 120000
  },
  "allMatches": [...],
  "totalMatches": 5
}
```

#### get_library_docs Response
```json
{
  "libraryID": "/vercel/next.js",
  "documentation": "# Next.js Documentation\n\nNext.js is a React framework...",
  "tokens": 10000,
  "actualTokens": 8500,
  "topic": "routing"
}
```

### Error Handling

Both tools implement comprehensive error handling:

- **Missing API Key**: Returns JSON error message
- **Invalid API Key**: Returns 401 error with helpful message
- **Library Not Found**: Returns 404 error with context
- **Network Issues**: Returns descriptive error messages
- **Empty Results**: Returns appropriate not found messages

### Configuration

#### Environment Variables
- `CONTEXT7_API_KEY`: Required API key for Context7 service

#### Dependencies
- `github.com/go-resty/resty/v2`: HTTP client library (already available in project)
- Standard Go libraries: `encoding/json`, `net/http`, `os`, `strings`, `strconv`

## Usage Examples

### Resolving a Library ID
```go
args := map[string]any{
    "libraryName": "react",
}
result, err := ResolveLibraryIDHandler(ctx, args)
```

### Fetching Documentation
```go
args := map[string]any{
    "context7CompatibleLibraryID": "/facebook/react",
    "tokens": 15000,
    "topic": "hooks",
}
result, err := GetLibraryDocsHandler(ctx, args)
```

## Testing

To test the implementation:

1. Set the `CONTEXT7_API_KEY` environment variable
2. Build and run the agent
3. Use the A2A protocol to call the tools with test parameters

## Compatibility

This implementation maintains full compatibility with:
- Context7 MCP server tool schemas
- Context7 API response formats
- ADL (Agent Definition Language) specifications
- A2A (Agent-to-Agent) protocol requirements

## Performance Considerations

- Direct API calls eliminate MCP protocol overhead
- HTTP client reuse for efficiency
- Minimal response processing and JSON marshaling
- Error responses avoid network calls when possible

## Security

- API key handled securely through environment variables
- Input validation on all parameters
- Proper HTTP status code handling
- No sensitive information logged or exposed