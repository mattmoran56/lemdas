# LEMDAS - Leeds Electron Microscope Data Access System

[![Open in Codespaces](https://classroom.github.com/assets/launch-codespace-7f7980b617ed060a017424585567c406b6ee15c891e84e1186181d67ecf80aa0.svg)](https://classroom.github.com/open-in-codespaces?assignment_repo_id=13179005)

## Overview

LEMDAS (Leeds Electron Microscope Data Access System) is a comprehensive data management platform developed in collaboration with the University of Leeds Electron Microscope team (LEMAS). The system provides academics and researchers with a centralized solution for storing, sharing, and viewing electron microscope data, facilitating collaborative research and data accessibility across the scientific community.

### Key Features

- **Secure Data Storage**: Store microscope images and datasets with metadata extraction
- **Collaborative Sharing**: Share datasets with individual users or research groups
- **Advanced Search**: Search across datasets and files using metadata attributes
- **File Processing**: Automatic metadata extraction from DM3 and TIF microscope formats
- **Access Control**: Fine-grained permissions for datasets and files
- **Preview Generation**: Automatic generation of image previews for quick viewing

## Data Model

```mermaid
erDiagram
    User ||--o{ Dataset : owns
    User ||--o{ File : owns
    User ||--o{ UserGroup : owns
    User ||--o{ StaredDataset : stars
    User ||--o{ Activity : performs
    User ||--o{ GroupMember : belongs_to
    User ||--o{ DatasetCollaborator : collaborates_on
    User ||--o{ UserShareDataset : receives_shared
    
    UserGroup ||--o{ GroupMember : has_members
    UserGroup ||--o{ GroupShareDataset : receives_shared_datasets
    
    Dataset ||--o{ File : contains
    Dataset ||--o{ DatasetAttribute : has_attributes
    Dataset ||--o{ DatasetCollaborator : has_collaborators
    Dataset ||--o{ UserShareDataset : shared_with_users
    Dataset ||--o{ GroupShareDataset : shared_with_groups
    Dataset ||--o{ StaredDataset : starred_by_users
    
    File ||--o{ FileAttribute : has_attributes
    File ||--o{ FileAttributeGroup : has_attribute_groups
    
    FileAttributeGroup ||--o{ FileAttribute : contains_attributes
    FileAttributeGroup ||--o{ FileAttributeGroup : has_children

    User {
        string id PK "UUID"
        string email "Unique email address"
        string firstName "User's first name"
        string lastName "User's last name"
        string avatar "Avatar image URL"
        string bio "User biography/description"
        int64 createdAt "Timestamp (milliseconds)"
        int64 updatedAt "Timestamp (milliseconds)"
    }
    
    Dataset {
        string id PK "UUID"
        string datasetName "Name of the dataset"
        string ownerId FK "References User.id"
        boolean isPublic "Public visibility flag"
        int64 createdAt "Timestamp (milliseconds)"
        int64 updatedAt "Timestamp (milliseconds)"
    }
    
    File {
        string id PK "UUID"
        string name "File name with extension"
        string ownerId FK "References User.id"
        string datasetId FK "References Dataset.id"
        string status "Status: uploaded|processing|processed|failed"
        int64 createdAt "Timestamp (milliseconds)"
        int64 updatedAt "Timestamp (milliseconds)"
    }
    
    DatasetAttribute {
        string id PK "UUID"
        string datasetId FK "References Dataset.id"
        string attributeName "Attribute key/name"
        string attributeValue "Attribute value"
        int64 createdAt "Timestamp (milliseconds)"
        int64 updatedAt "Timestamp (milliseconds)"
    }
    
    FileAttribute {
        string id PK "UUID"
        string fileId FK "References File.id"
        string attributeName "Attribute key/name"
        string attributeValue "Attribute value"
        string attributeGroupId FK "References FileAttributeGroup.id (nullable)"
        int64 createdAt "Timestamp (milliseconds)"
        int64 updatedAt "Timestamp (milliseconds)"
    }
    
    FileAttributeGroup {
        string id PK "UUID"
        string attributeGroupName "Group name"
        string fileId FK "References File.id"
        string parentGroupId FK "Self-reference for hierarchy (nullable)"
        int64 createdAt "Timestamp (milliseconds)"
        int64 updatedAt "Timestamp (milliseconds)"
    }
    
    UserGroup {
        string id PK "UUID"
        string groupName "Name of the group"
        string ownerId FK "References User.id"
        int64 createdAt "Timestamp (milliseconds)"
        int64 updatedAt "Timestamp (milliseconds)"
    }
    
    GroupMember {
        string id PK "UUID"
        string groupId FK "References UserGroup.id"
        string userId FK "References User.id"
        int64 createdAt "Timestamp (milliseconds)"
        int64 updatedAt "Timestamp (milliseconds)"
    }
    
    DatasetCollaborator {
        string id PK "UUID"
        string userId FK "References User.id"
        string datasetId FK "References Dataset.id"
        int64 createdAt "Timestamp (milliseconds)"
        int64 updatedAt "Timestamp (milliseconds)"
    }
    
    UserShareDataset {
        string id PK "UUID"
        string userId FK "References User.id"
        string datasetId FK "References Dataset.id"
        boolean writeAccess "Write permission flag"
        int64 createdAt "Timestamp (milliseconds)"
        int64 updatedAt "Timestamp (milliseconds)"
    }
    
    GroupShareDataset {
        string id PK "UUID"
        string groupId FK "References UserGroup.id"
        string datasetId FK "References Dataset.id"
        boolean writeAccess "Write permission flag"
        int64 createdAt "Timestamp (milliseconds)"
        int64 updatedAt "Timestamp (milliseconds)"
    }
    
    StaredDataset {
        string id PK "UUID"
        string userId FK "References User.id"
        string datasetId FK "References Dataset.id"
        int64 createdAt "Timestamp (milliseconds)"
        int64 updatedAt "Timestamp (milliseconds)"
    }
    
    Activity {
        string id PK "UUID"
        string type "Activity type"
        string userId FK "References User.id"
        string details "Activity details JSON"
        int64 createdAt "Timestamp (milliseconds)"
        int64 updatedAt "Timestamp (milliseconds)"
    }
```

### Data Model Details

#### Core Entities

- **User**: Central entity representing system users with profile information (email, name, avatar, bio)
- **Dataset**: Collection of related files owned by a user, can be public or private
- **File**: Individual files within datasets, tracked with processing status
- **UserGroup**: Named groups for collaborative access management

#### Metadata & Attributes

- **DatasetAttribute**: Key-value pairs for dataset-level metadata
- **FileAttribute**: Key-value pairs for file-level metadata, can be grouped
- **FileAttributeGroup**: Hierarchical grouping of file attributes (supports nested groups via parentGroupId)

#### Access Control & Sharing

- **DatasetCollaborator**: Users with direct collaboration rights on a dataset
- **UserShareDataset**: Individual user sharing with read/write permissions
- **GroupShareDataset**: Group-based sharing with read/write permissions
- **GroupMember**: Membership tracking for user groups

#### User Interaction

- **StaredDataset**: User bookmarks/favorites for quick access
- **Activity**: Audit log of user actions in the system

#### Key Design Patterns

1. **UUID Primary Keys**: All entities use string UUIDs for distributed system compatibility
2. **Soft Timestamps**: CreatedAt/UpdatedAt stored as int64 milliseconds for precision
3. **Hierarchical Attributes**: FileAttributeGroup supports nested structure for complex metadata
4. **Flexible Permissions**: Separate collaborator and sharing models for granular access control
5. **Status Tracking**: Files have status field for processing pipeline (uploaded → processing → processed)

## Quick Start

### Prerequisites

- Node.js (v18+)
- Go (v1.21+)
- Python (v3.12+) with Poetry
- MySQL (v8.0+)
- Azure Storage Account (for blob storage)

### Environment Setup

Create `.env` files for each service with required variables:

```bash
# Database Configuration
DB_USERNAME=your_db_user
DB_PASSWORD=your_db_password
DB_NAME=lemdas
DB_HOST=localhost
DB_PORT=3306

# JWT Secret
JWT_SECRET=your_jwt_secret

# Azure Storage
AZURE_STORAGE_ACCOUNT=your_storage_account
AZURE_STORAGE_KEY=your_storage_key
```

### Running the Services

1. **Start MySQL Database**
   ```bash
   mysql.server start  # or use Docker
   ```

2. **Run Go Microservices**
   ```bash
   cd internal
   go mod download
   
   # Run each service in separate terminals
   go run auth/main.go
   go run webApi/main.go
   go run search/main.go
   go run upload/main.go
   ```

3. **Run Python File Processor**
   ```bash
   cd file_processor/fileProcessor
   poetry install
   poetry run uvicorn main:app --reload
   ```

4. **Run React Frontend**
   ```bash
   cd frontend
   npm install
   npm start
   ```

The application will be available at `http://localhost:3000`

## Architecture

LEMDAS uses a microservices architecture with:
- **Frontend**: React SPA for user interface
- **Auth Service**: JWT token management
- **Web API**: Core CRUD operations
- **Search Service**: Advanced search capabilities
- **Upload Service**: File upload/download with Azure Blob Storage
- **File Processor**: Python service for metadata extraction from microscope formats

## License

This project was developed as part of an academic collaboration with the University of Leeds.