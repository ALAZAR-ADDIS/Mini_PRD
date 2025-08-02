// MongoDB Database Initialization Script for Blog Platform
// Run this script in MongoDB Atlas shell

// Switch to blog_platform database
use blog_platform

print("🚀 Initializing Blog Platform Database...")

// Create collections (if they don't exist)
const collections = [
  "users",
  "blog_posts", 
  "comments",
  "user_interactions",
  "auth_tokens",
  "tags",
  "categories",
  "ai_suggestions",
  "user_sessions"
]

collections.forEach(collection => {
  if (!db.getCollectionNames().includes(collection)) {
    db.createCollection(collection)
    print(`✅ Created collection: ${collection}`)
  } else {
    print(`ℹ️  Collection already exists: ${collection}`)
  }
})

print("\n📊 Creating indexes...")

// Users Collection Indexes
try {
  db.users.createIndex({ "email": 1 }, { unique: true })
  print("✅ Users email index created")
} catch (e) {
  print("ℹ️  Users email index already exists")
}

try {
  db.users.createIndex({ "username": 1 }, { unique: true })
  print("✅ Users username index created")
} catch (e) {
  print("ℹ️  Users username index already exists")
}

try {
  db.users.createIndex({ "role": 1 })
  print("✅ Users role index created")
} catch (e) {
  print("ℹ️  Users role index already exists")
}

try {
  db.users.createIndex({ "isActive": 1 })
  print("✅ Users isActive index created")
} catch (e) {
  print("ℹ️  Users isActive index already exists")
}

// Blog Posts Collection Indexes
try {
  db.blog_posts.createIndex({ "authorId": 1 })
  print("✅ Blog posts authorId index created")
} catch (e) {
  print("ℹ️  Blog posts authorId index already exists")
}

try {
  db.blog_posts.createIndex({ "status": 1 })
  print("✅ Blog posts status index created")
} catch (e) {
  print("ℹ️  Blog posts status index already exists")
}

try {
  db.blog_posts.createIndex({ "publishedAt": -1 })
  print("✅ Blog posts publishedAt index created")
} catch (e) {
  print("ℹ️  Blog posts publishedAt index already exists")
}

try {
  db.blog_posts.createIndex({ "viewCount": -1 })
  print("✅ Blog posts viewCount index created")
} catch (e) {
  print("ℹ️  Blog posts viewCount index already exists")
}

try {
  db.blog_posts.createIndex({ "likeCount": -1 })
  print("✅ Blog posts likeCount index created")
} catch (e) {
  print("ℹ️  Blog posts likeCount index already exists")
}

try {
  db.blog_posts.createIndex({ "tags": 1 })
  print("✅ Blog posts tags index created")
} catch (e) {
  print("ℹ️  Blog posts tags index already exists")
}

try {
  db.blog_posts.createIndex({ "category": 1 })
  print("✅ Blog posts category index created")
} catch (e) {
  print("ℹ️  Blog posts category index already exists")
}

try {
  db.blog_posts.createIndex({ "slug": 1 }, { unique: true })
  print("✅ Blog posts slug index created")
} catch (e) {
  print("ℹ️  Blog posts slug index already exists")
}

try {
  db.blog_posts.createIndex({
    "title": "text",
    "content": "text",
    "tags": "text"
  })
  print("✅ Blog posts text search index created")
} catch (e) {
  print("ℹ️  Blog posts text search index already exists")
}

// Comments Collection Indexes
try {
  db.comments.createIndex({ "postId": 1 })
  print("✅ Comments postId index created")
} catch (e) {
  print("ℹ️  Comments postId index already exists")
}

try {
  db.comments.createIndex({ "authorId": 1 })
  print("✅ Comments authorId index created")
} catch (e) {
  print("ℹ️  Comments authorId index already exists")
}

try {
  db.comments.createIndex({ "parentCommentId": 1 })
  print("✅ Comments parentCommentId index created")
} catch (e) {
  print("ℹ️  Comments parentCommentId index already exists")
}

try {
  db.comments.createIndex({ "createdAt": -1 })
  print("✅ Comments createdAt index created")
} catch (e) {
  print("ℹ️  Comments createdAt index already exists")
}

// User Interactions Collection Indexes
try {
  db.user_interactions.createIndex(
    { "userId": 1, "postId": 1, "interactionType": 1 },
    { unique: true }
  )
  print("✅ User interactions compound index created")
} catch (e) {
  print("ℹ️  User interactions compound index already exists")
}

try {
  db.user_interactions.createIndex({ "postId": 1, "interactionType": 1 })
  print("✅ User interactions postId+type index created")
} catch (e) {
  print("ℹ️  User interactions postId+type index already exists")
}

try {
  db.user_interactions.createIndex({ "userId": 1 })
  print("✅ User interactions userId index created")
} catch (e) {
  print("ℹ️  User interactions userId index already exists")
}

// Auth Tokens Collection Indexes
try {
  db.auth_tokens.createIndex({ "userId": 1, "tokenType": 1 })
  print("✅ Auth tokens userId+type index created")
} catch (e) {
  print("ℹ️  Auth tokens userId+type index already exists")
}

try {
  db.auth_tokens.createIndex({ "token": 1 }, { unique: true })
  print("✅ Auth tokens token index created")
} catch (e) {
  print("ℹ️  Auth tokens token index already exists")
}

try {
  db.auth_tokens.createIndex({ "expiresAt": 1 }, { expireAfterSeconds: 0 })
  print("✅ Auth tokens TTL index created")
} catch (e) {
  print("ℹ️  Auth tokens TTL index already exists")
}

// Tags Collection Indexes
try {
  db.tags.createIndex({ "name": 1 }, { unique: true })
  print("✅ Tags name index created")
} catch (e) {
  print("ℹ️  Tags name index already exists")
}

try {
  db.tags.createIndex({ "postCount": -1 })
  print("✅ Tags postCount index created")
} catch (e) {
  print("ℹ️  Tags postCount index already exists")
}

// Categories Collection Indexes
try {
  db.categories.createIndex({ "name": 1 }, { unique: true })
  print("✅ Categories name index created")
} catch (e) {
  print("ℹ️  Categories name index already exists")
}

try {
  db.categories.createIndex({ "slug": 1 }, { unique: true })
  print("✅ Categories slug index created")
} catch (e) {
  print("ℹ️  Categories slug index already exists")
}

try {
  db.categories.createIndex({ "postCount": -1 })
  print("✅ Categories postCount index created")
} catch (e) {
  print("ℹ️  Categories postCount index already exists")
}

// AI Suggestions Collection Indexes
try {
  db.ai_suggestions.createIndex({ "userId": 1 })
  print("✅ AI suggestions userId index created")
} catch (e) {
  print("ℹ️  AI suggestions userId index already exists")
}

try {
  db.ai_suggestions.createIndex({ "postId": 1 })
  print("✅ AI suggestions postId index created")
} catch (e) {
  print("ℹ️  AI suggestions postId index already exists")
}

try {
  db.ai_suggestions.createIndex({ "suggestionType": 1 })
  print("✅ AI suggestions type index created")
} catch (e) {
  print("ℹ️  AI suggestions type index already exists")
}

try {
  db.ai_suggestions.createIndex({ "isUsed": 1 })
  print("✅ AI suggestions isUsed index created")
} catch (e) {
  print("ℹ️  AI suggestions isUsed index already exists")
}

// User Sessions Collection Indexes
try {
  db.user_sessions.createIndex({ "userId": 1 })
  print("✅ User sessions userId index created")
} catch (e) {
  print("ℹ️  User sessions userId index already exists")
}

try {
  db.user_sessions.createIndex({ "sessionId": 1 }, { unique: true })
  print("✅ User sessions sessionId index created")
} catch (e) {
  print("ℹ️  User sessions sessionId index already exists")
}

try {
  db.user_sessions.createIndex({ "expiresAt": 1 }, { expireAfterSeconds: 0 })
  print("✅ User sessions TTL index created")
} catch (e) {
  print("ℹ️  User sessions TTL index already exists")
}

print("\n📝 Inserting initial data...")

// Insert default categories if they don't exist
const defaultCategories = [
  {
    name: "Technology",
    description: "Technology related articles",
    slug: "technology",
    postCount: 0,
    createdAt: new Date()
  },
  {
    name: "Lifestyle",
    description: "Lifestyle and personal development",
    slug: "lifestyle",
    postCount: 0,
    createdAt: new Date()
  },
  {
    name: "Business",
    description: "Business and entrepreneurship",
    slug: "business",
    postCount: 0,
    createdAt: new Date()
  }
]

defaultCategories.forEach(category => {
  try {
    db.categories.insertOne(category)
    print(`✅ Inserted category: ${category.name}`)
  } catch (e) {
    print(`ℹ️  Category already exists: ${category.name}`)
  }
})

// Insert default tags if they don't exist
const defaultTags = [
  {
    name: "golang",
    description: "Go programming language",
    postCount: 0,
    createdAt: new Date()
  },
  {
    name: "mongodb",
    description: "MongoDB database",
    postCount: 0,
    createdAt: new Date()
  },
  {
    name: "api",
    description: "Application Programming Interface",
    postCount: 0,
    createdAt: new Date()
  }
]

defaultTags.forEach(tag => {
  try {
    db.tags.insertOne(tag)
    print(`✅ Inserted tag: ${tag.name}`)
  } catch (e) {
    print(`ℹ️  Tag already exists: ${tag.name}`)
  }
})

print("\n🎉 Database initialization completed successfully!")
print("📊 Summary:")
print(`- Collections: ${collections.length}`)
print(`- Default categories: ${defaultCategories.length}`)
print(`- Default tags: ${defaultTags.length}`)
print("\n🚀 Your blog platform database is ready!") 