# Web Tech Solutions - Excel File Handling API

## Overview

This API allows you to upload Excel files, extract their data, and process it for analysis.

## Endpoints

### POST /upload

- **Description**: Upload an Excel file to process its data.
- **Request**:
  - `form-data` with the key `file` (the Excel file to upload).
- **Response**:
  - `200 OK` if successful.
  - `400 Bad Request` for invalid file formats.

## Example Request
bash
curl -X POST http://localhost:8080/upload -F "file=@/path/to/excel.xlsx"
## Libraries Used

- **Gin**: HTTP web framework for Go.
- **Excelize**: Library for reading and writing Excel files.


---
