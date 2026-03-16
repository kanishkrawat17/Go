# Frontend System Design: Chat Application (Full Interview Guide)

> **Context:** You're in a 45-60 minute frontend system design round. The interviewer says:
> *"Design a chat application like WhatsApp Web or Slack."*
>
> This guide walks you through **exactly what to say, in what order, and why**.
> Written for someone with basic frontend knowledge -- no hand-waving.

---

## Table of Contents

1. [Before We Start: How This Round Works](#before-we-start-how-this-round-works)
2. [Step 1: Requirements Clarification (Minutes 0-5)](#step-1-requirements-clarification-minutes-0-5)
3. [Step 2: High-Level Architecture (Minutes 5-10)](#step-2-high-level-architecture-minutes-5-10)
4. [Step 3: API Endpoints -- REST + WebSocket Events (Minutes 10-15)](#step-3-api-endpoints----rest--websocket-events-minutes-10-15)
5. [Step 4: Data Model Mapped to UI (Minutes 15-22)](#step-4-data-model-mapped-to-ui-minutes-15-22)
6. [Step 5: Component Design (Minutes 22-27)](#step-5-component-design-minutes-22-27)
7. [Step 6: WebSocket Connection Management (Minutes 27-35)](#step-6-websocket-connection-management-minutes-27-35)
8. [Step 7: Message Ordering & Delivery Receipts (Minutes 35-42)](#step-7-message-ordering--delivery-receipts-minutes-35-42)
9. [Step 8: Offline Message Queue (Minutes 42-47)](#step-8-offline-message-queue-minutes-42-47)
10. [Step 9: Scroll Position & Infinite Scroll Up (Minutes 47-53)](#step-9-scroll-position--infinite-scroll-up-minutes-47-53)
11. [Step 10: Wrap-Up & Tradeoffs (Minutes 53-58)](#step-10-wrap-up--tradeoffs-minutes-53-58)
12. [Cross Questions Bank (with Answers)](#cross-questions-bank-with-answers)
13. [Cheat Sheet: Things That Impress Interviewers](#cheat-sheet-things-that-impress-interviewers)

---

## Before We Start: How This Round Works

### What the interviewer is REALLY evaluating:

| Signal | What They Look For | Weight |
|--------|-------------------|--------|
| **Structured thinking** | Do you have a framework? Do you go step-by-step? | Very High |
| **Requirement gathering** | Do you ask questions or just start designing? | High |
| **Trade-off awareness** | Can you explain WHY you chose X over Y? | Very High |
| **Depth in key areas** | Can you go deep when probed? | High |
| **Communication** | Do you think aloud? Can you draw/explain clearly? | High |

### What they are NOT evaluating:
- Memorized system designs
- Perfect answers
- Knowing every API by name

### Golden rule:
**Talk like you're designing this with a teammate, not reciting an answer to a teacher.**

Say things like:
- "One approach would be... but the trade-off is..."
- "I'd lean towards X because..."
- "Let me think about the edge cases here..."

---

## Step 1: Requirements Clarification (Minutes 0-5)

### What to say (literally):

> "Before I start designing, I'd like to clarify a few things to make sure I'm solving the right problem."

Then ask these questions. **Do not skip this.** Jumping straight into design is the #1 mistake.

### Functional Requirements (ask the interviewer):

| Question | Why You're Asking | Likely Answer |
|----------|------------------|---------------|
| "Is this 1-on-1 chat, group chat, or both?" | Scope matters hugely -- group chat is 3x more complex | "Focus on 1-on-1, mention how you'd extend to group" |
| "Do we need real-time messaging or is near-real-time okay?" | Determines WebSocket vs polling | "Real-time" |
| "What message types? Just text, or also images/files/voice?" | Affects data model and upload flow | "Start with text, mention media" |
| "Do we need read receipts / delivery status?" | Adds complexity to message state | "Yes" |
| "Is there a typing indicator?" | Real-time feature, needs thought | "Yes, nice to have" |
| "Do we need message search?" | Completely different problem if yes | "Not in scope for now" |
| "Do we need offline support?" | Affects architecture significantly | "Yes" |

### Non-Functional Requirements (state these yourself):

> "For non-functional requirements, I'll assume:
> - **Low latency**: messages should appear in < 200ms under normal conditions
> - **Reliability**: no messages should be lost, even on flaky networks
> - **Message ordering**: messages must appear in the correct order
> - **Offline resilience**: users should be able to queue messages when offline
> - **Scalable**: should handle a conversation with thousands of messages (scroll performance)"

### After clarification, summarize:

> "So I'm designing a 1-on-1 chat app with real-time messaging, delivery/read receipts, typing indicators, offline support, and text messages. Let me start with the high-level architecture."

---

## Step 2: High-Level Architecture (Minutes 5-10)

### Draw/describe this diagram:

```
┌──────────────────────────────────────────────────────┐
│                    BROWSER (Client)                   │
│                                                       │
│  ┌─────────────┐  ┌──────────────┐  ┌──────────────┐ │
│  │  React App   │  │  WebSocket   │  │   Service    │ │
│  │  (UI Layer)  │◄─┤  Manager     │  │   Worker     │ │
│  │              │  │  (real-time)  │  │  (offline)   │ │
│  └──────┬───────┘  └──────┬───────┘  └──────┬───────┘ │
│         │                 │                  │         │
│  ┌──────▼─────────────────▼──────────────────▼──────┐ │
│  │              State Management (Zustand)            │ │
│  └──────────────────────┬────────────────────────────┘ │
│                         │                              │
│  ┌──────────────────────▼────────────────────────────┐ │
│  │                   IndexedDB                        │ │
│  │            (offline cache, message queue)          │ │
│  └───────────────────────────────────────────────────┘ │
└──────────────────────────┬───────────────────────────┘
                           │
              ┌────────────▼────────────┐
              │     WebSocket Server    │ ← persistent connection (real-time)
              └────────────┬────────────┘
              ┌────────────▼────────────┐
              │      REST API Server    │ ← request-response (history, auth)
              └────────────┬────────────┘
              ┌────────────▼────────────┐
              │        Database         │
              └─────────────────────────┘
```

### What to say:

> "I see two communication channels between client and server:
>
> 1. **WebSocket** -- for real-time stuff: message delivery, typing indicators, online/offline status, read receipts. This is a persistent, bi-directional connection.
>
> 2. **REST API** -- for request-response stuff: login, fetching message history (with pagination), user profile. These are one-off requests, WebSocket would be overkill.
>
> On the client, a **Zustand store** is the single source of truth. The UI reads from it, WebSocket events write to it, and it syncs to **IndexedDB** for offline support."

### Cross-question you might get here:

> **Interviewer: "Why WebSocket and not SSE or long polling?"**

| Option | How it works | Good for chat? |
|--------|-------------|----------------|
| **Long Polling** | Client keeps making HTTP requests, server holds each open until data arrives | Works but wasteful -- each poll = new connection, high latency |
| **SSE (Server-Sent Events)** | Server pushes to client over one HTTP connection | **One-way only** (server → client). Client still needs REST to send messages. Not ideal. |
| **WebSocket** | Persistent two-way pipe. Both sides send whenever they want. | Perfect for chat. Low overhead after initial handshake. |

> "I'd use **WebSocket as primary**, with **long polling as fallback** for corporate firewalls that block WebSocket."

---

## Step 3: API Endpoints -- REST + WebSocket Events (Minutes 10-15)

### What to say:

> "Let me define the contract between frontend and backend -- what REST endpoints I need and what WebSocket events flow in each direction."

### REST API Endpoints

| Method | Endpoint | What it does | When it's called |
|--------|----------|-------------|-----------------|
| `POST` | `/auth/login` | Login, get access token | App open |
| `GET` | `/conversations` | Get user's conversation list | App open, pull-to-refresh |
| `GET` | `/conversations/:id/messages?before=timestamp&limit=30` | Get older messages (paginated) | Scroll up to load history |
| `POST` | `/conversations` | Create a new conversation | User starts a new chat |
| `POST` | `/media/upload` | Upload image/file, get back URL | User attaches media |
| `GET` | `/users/:id` | Get user profile info | View profile |

#### Key design choices to explain:

> **Pagination:** "I use **cursor-based pagination** (`before=timestamp`) not page numbers (`page=3`). Why? Because in a chat app, new messages are being added constantly. If I'm on page 3 and 10 new messages arrive, page 3 now shows different messages. With cursor-based, I say 'give me 30 messages older than this timestamp' which is stable regardless of new messages."

> **Why REST for history, not WebSocket?** "Message history is a classic request-response pattern -- I ask, server answers. REST gives me HTTP caching (browser can cache responses), retry semantics, and it's simpler. WebSocket is for stuff that happens *without* the client asking -- incoming messages, typing events, etc."

### WebSocket Events

#### Client → Server (what the frontend SENDS):

| Event Type | Payload | When it's sent |
|-----------|---------|---------------|
| `send_message` | `{id, conversationId, content, type, clientTimestamp}` | User hits Send |
| `typing` | `{conversationId}` | User is typing (throttled, every 3s) |
| `read_ack` | `{conversationId, lastReadMessageId}` | User views a conversation |
| `delivery_ack` | `{messageId}` | Client receives a new message |
| `ping` | `{}` | Every 30s (heartbeat) |
| `sync` | `{lastReceivedTimestamp}` | On reconnect (get missed messages) |

#### Server → Client (what the frontend RECEIVES):

| Event Type | Payload | What the UI does |
|-----------|---------|-----------------|
| `new_message` | `{message: Message}` | Add to message list, update conversation preview |
| `message_ack` | `{id, timestamp, status: 'sent'}` | Update message from 'sending' → 'sent' (show ✓) |
| `delivery_receipt` | `{messageId}` | Update message to 'delivered' (show ✓✓) |
| `read_receipt` | `{conversationId, lastReadMessageId}` | Update messages up to that ID to 'read' (show blue ✓✓) |
| `typing` | `{conversationId, userId}` | Show "Alice is typing..." |
| `presence` | `{userId, status: 'online'/'offline'}` | Update green dot on avatar |
| `pong` | `{}` | Connection is alive (heartbeat response) |

### The full picture at a glance:

```
┌───────────────────────────────────────────────────────────────┐
│                         REST API                              │
│                                                               │
│  App loads ──► GET /conversations ──► Show conversation list  │
│  Scroll up ──► GET /conversations/:id/messages ──► Prepend    │
│  Login    ──► POST /auth/login ──► Get token                  │
│                                                               │
├───────────────────────────────────────────────────────────────┤
│                     WebSocket (always open)                    │
│                                                               │
│  Client ──► send_message, typing, read_ack, ping              │
│  Server ──► new_message, message_ack, delivery_receipt,       │
│             read_receipt, typing, presence, pong               │
└───────────────────────────────────────────────────────────────┘
```

---

## Step 4: Data Model Mapped to UI (Minutes 15-22)

### What to say:

> "Let me show the data model by mapping it directly to the UI. Each piece of state exists because a specific part of the screen needs it."

### The UI and What Data Powers Each Part

```
┌─────────────────────────────────────────────────────────────────────┐
│                     connectionStatus ◄────────────────┐             │
│  ┌──────────────────┐              ┌──────────────────┤             │
│  │ "Reconnecting..."│  (shown when │  "Online" |      │             │
│  │ "Connecting..."  │  disconnected│  "Connecting" |  │             │
│  └──────────────────┘   only)      │  "Disconnected"  │             │
│                                    └──────────────────┘             │
│                                                                     │
│  ┌──────────────────────────┐  ┌────────────────────────────────┐  │
│  │      SIDEBAR             │  │         CHAT WINDOW             │  │
│  │                          │  │                                  │  │
│  │  ┌─ SearchBar ─────────┐│  │  ┌─ ChatHeader ───────────────┐ │  │
│  │  │ [Search messages...] ││  │  │                             │ │  │
│  │  └──────────────────────┘│  │  │  participant.name           │ │  │
│  │                          │  │  │  onlineUsers.has(id)        │ │  │
│  │  ┌─ ConversationItem ──┐│  │  │    → "Online" / "Offline"   │ │  │
│  │  │                     ││  │  │                             │ │  │
│  │  │  (○) Name1  Status  ││  │  │  typingIndicators[convId]   │ │  │
│  │  │                     ││  │  │    → "Alice is typing..."    │ │  │
│  │  │ ┌─ DATA USED: ────┐││  │  └─────────────────────────────┘ │  │
│  │  │ │ conversation.    │││  │                                  │  │
│  │  │ │  participants[0] │││  │  ┌─ MessageList ───────────────┐ │  │
│  │  │ │   → name, avatar │││  │  │                             │ │  │
│  │  │ │ conversation.    │││  │  │  messages[conversationId]   │ │  │
│  │  │ │  lastMessage     │││  │  │   → array of Message objects│ │  │
│  │  │ │   → preview text │││  │  │                             │ │  │
│  │  │ │ conversation.    │││  │  │  ┌─ MessageBubble ────────┐ │ │  │
│  │  │ │  unreadCount     │││  │  │  │                        │ │ │  │
│  │  │ │   → badge "3"   │││  │  │  │  message.content       │ │ │  │
│  │  │ │ onlineUsers      │││  │  │  │   → "Hey, how are you?"│ │ │  │
│  │  │ │  .has(oderId)    │││  │  │  │                        │ │ │  │
│  │  │ │   → green dot    │││  │  │  │  message.timestamp     │ │ │  │
│  │  │ └─────────────────┘││  │  │  │   → "2:34 PM"          │ │ │  │
│  │  └─────────────────────┘│  │  │  │                        │ │ │  │
│  │                          │  │  │  │  message.status        │ │ │  │
│  │  ┌─ ConversationItem ──┐│  │  │  │   → ✓ / ✓✓ / ✓✓ blue │ │ │  │
│  │  │ (○) Name2  Status   ││  │  │  └────────────────────────┘ │ │  │
│  │  └─────────────────────┘│  │  │                             │ │  │
│  │                          │  │  └─────────────────────────────┘ │  │
│  │  ┌─ ConversationItem ──┐│  │                                  │  │
│  │  │ (○) Name3  Status   ││  │  ┌─ MessageInput ─────────────┐ │  │
│  │  └─────────────────────┘│  │  │                             │ │  │
│  │                          │  │  │ [+] [Type your message...] │ │  │
│  │  ┌─ ConversationItem ──┐│  │  │                     [Send] │ │  │
│  │  │ (○) Name4  Status   ││  │  │                             │ │  │
│  │  └─────────────────────┘│  │  └─────────────────────────────┘ │  │
│  └──────────────────────────┘  └────────────────────────────────┘  │
└─────────────────────────────────────────────────────────────────────┘
```

### Now the data model -- each field is justified by the UI above:

```typescript
interface ChatState {

  // ── Powers the SIDEBAR (conversation list) ──────────────────
  conversations: Record<string, Conversation>
  //  Why Record (hash map) and not array?
  //  → When a new message arrives, I need to update ONE conversation's
  //    lastMessage. With an array I'd search through all of them (slow).
  //    With Record, it's conversations["conv_123"] → instant.

  // ── Powers the CHAT WINDOW (message area) ───────────────────
  messages: Record<string, Message[]>
  //  Key = conversationId, Value = sorted array of messages
  //  → I only load messages for the conversation the user is viewing.
  //    Switching conversations = just read a different key. No re-fetching.

  activeConversationId: string | null
  //  → Which conversation is currently open in the Chat Window.
  //    null = no chat selected (show "Select a chat to start" screen).

  // ── Powers the GREEN DOT on avatars ─────────────────────────
  onlineUsers: Set<string>
  //  Set of user IDs who are currently online.
  //  → Set gives O(1) lookup: onlineUsers.has("user_456") → true/false
  //  → Updated via WebSocket 'presence' events.

  // ── Powers "Alice is typing..." in ChatHeader ───────────────
  typingIndicators: Record<string, string[]>
  //  Key = conversationId, Value = array of userIds who are typing
  //  → typingIndicators["conv_123"] = ["user_456"]
  //  → Purely in-memory. Never saved to database or IndexedDB.
  //  → Auto-clears after 4 seconds of no new typing event.

  // ── Powers the "Reconnecting..." banner at top ──────────────
  connectionStatus: 'connected' | 'connecting' | 'disconnected'
  //  → Shows/hides the connection banner.
  //  → Updated by the WebSocket Manager.

  // ── Powers the offline queue (invisible to user) ────────────
  pendingMessages: Message[]
  //  → Messages sent while offline. Stored in IndexedDB too.
  //  → Flushed to server when connection is restored.
}
```

### The two core data objects:

```typescript
interface Conversation {
  id: string
  participants: string[]           // → name + avatar in sidebar
  lastMessage: Message | null      // → preview text in sidebar
  unreadCount: number              // → badge "3" in sidebar
  updatedAt: number                // → sort conversations (most recent first)
}
```

```typescript
interface Message {
  id: string                       // UUID generated on client (for optimistic UI)
  conversationId: string           // which conversation this belongs to
  senderId: string                 // who sent it → left/right bubble alignment
  content: string                  // the actual text
  type: 'text' | 'image' | 'file' // for future extensibility

  timestamp: number                // SERVER assigns this → used for ordering
  clientTimestamp: number          // CLIENT sets this → used for optimistic display

  status: 'sending'                // ⏳ gray clock (just hit send, waiting for server)
        | 'sent'                   // ✓  single gray check (server received it)
        | 'delivered'              // ✓✓ double gray check (other user's device got it)
        | 'read'                   // ✓✓ double blue check (other user opened the chat)
        | 'failed'                 // ❌ red icon + retry button (server never ACKed)
}
```

### Mock Data: What the state ACTUALLY looks like (with real values)

Here's the entire `ChatState` filled with mock data. Below it, I show exactly what the user sees rendered from each value.

```typescript
const chatState: ChatState = {

  // ═══════════════════════════════════════════════════════════
  //  CONVERSATIONS (powers the sidebar)
  // ═══════════════════════════════════════════════════════════
  conversations: {
    "conv_101": {
      id: "conv_101",
      participants: ["user_me", "user_alice"],
      lastMessage: {
        id: "msg_550",
        conversationId: "conv_101",
        senderId: "user_alice",
        content: "Sure, let's meet at 6!",
        type: "text",
        timestamp: 1710421200000,    // Mar 14, 2026 5:00 PM
        clientTimestamp: 1710421200000,
        status: "read",
      },
      unreadCount: 0,
      updatedAt: 1710421200000,
    },

    "conv_102": {
      id: "conv_102",
      participants: ["user_me", "user_bob"],
      lastMessage: {
        id: "msg_620",
        conversationId: "conv_102",
        senderId: "user_bob",
        content: "Can you review my PR?",
        type: "text",
        timestamp: 1710418800000,    // Mar 14, 2026 4:20 PM
        clientTimestamp: 1710418800000,
        status: "delivered",
      },
      unreadCount: 3,
      updatedAt: 1710418800000,
    },

    "conv_103": {
      id: "conv_103",
      participants: ["user_me", "user_carol"],
      lastMessage: null,              // New conversation, no messages yet
      unreadCount: 0,
      updatedAt: 1710410000000,
    },
  },

  // ═══════════════════════════════════════════════════════════
  //  MESSAGES (powers the chat window for each conversation)
  // ═══════════════════════════════════════════════════════════
  messages: {
    // Messages for the conversation with Alice (conv_101)
    "conv_101": [
      {
        id: "msg_501",
        conversationId: "conv_101",
        senderId: "user_me",         // I sent this → RIGHT side bubble
        content: "Hey Alice!",
        type: "text",
        timestamp: 1710410000000,    // Mar 14, 2026 1:53 PM
        clientTimestamp: 1710410000000,
        status: "read",              // ✓✓ blue
      },
      {
        id: "msg_502",
        conversationId: "conv_101",
        senderId: "user_alice",      // Alice sent this → LEFT side bubble
        content: "Hi! What's up?",
        type: "text",
        timestamp: 1710410060000,    // 1 min later
        clientTimestamp: 1710410060000,
        status: "read",              // (status only matters for MY messages)
      },
      {
        id: "msg_503",
        conversationId: "conv_101",
        senderId: "user_me",
        content: "Wanna grab dinner tonight?",
        type: "text",
        timestamp: 1710420000000,    // Mar 14, 2026 4:40 PM
        clientTimestamp: 1710420000000,
        status: "read",              // ✓✓ blue
      },
      {
        id: "msg_550",
        conversationId: "conv_101",
        senderId: "user_alice",
        content: "Sure, let's meet at 6!",
        type: "text",
        timestamp: 1710421200000,    // Mar 14, 2026 5:00 PM
        clientTimestamp: 1710421200000,
        status: "read",
      },
    ],

    // Messages for the conversation with Bob (conv_102)
    "conv_102": [
      {
        id: "msg_610",
        conversationId: "conv_102",
        senderId: "user_me",
        content: "How's the sprint going?",
        type: "text",
        timestamp: 1710415000000,
        clientTimestamp: 1710415000000,
        status: "read",              // ✓✓ blue
      },
      {
        id: "msg_615",
        conversationId: "conv_102",
        senderId: "user_bob",
        content: "Almost done, just one more ticket",
        type: "text",
        timestamp: 1710416000000,
        clientTimestamp: 1710416000000,
        status: "delivered",
      },
      {
        id: "msg_618",
        conversationId: "conv_102",
        senderId: "user_me",
        content: "Nice, let me know if you need help",
        type: "text",
        timestamp: 1710417000000,
        clientTimestamp: 1710417000000,
        status: "delivered",         // ✓✓ gray (Bob's device got it but he hasn't opened chat)
      },
      {
        id: "msg_620",
        conversationId: "conv_102",
        senderId: "user_bob",
        content: "Can you review my PR?",
        type: "text",
        timestamp: 1710418800000,
        clientTimestamp: 1710418800000,
        status: "delivered",
      },
      // This message is still sending (I just typed it while offline)
      {
        id: "msg_625",
        conversationId: "conv_102",
        senderId: "user_me",
        content: "Sure, sending the link?",
        type: "text",
        timestamp: 1710419000000,
        clientTimestamp: 1710419000000,
        status: "sending",           // ⏳ clock icon (waiting for server ACK)
      },
    ],

    // conv_103 has no messages yet (Carol, new conversation)
    "conv_103": [],
  },

  // ═══════════════════════════════════════════════════════════
  //  ACTIVE CONVERSATION
  // ═══════════════════════════════════════════════════════════
  activeConversationId: "conv_101",  // Alice's chat is currently open

  // ═══════════════════════════════════════════════════════════
  //  ONLINE STATUS (powers the green dots)
  // ═══════════════════════════════════════════════════════════
  onlineUsers: new Set(["user_alice", "user_carol"]),
  //  Alice  → green dot ●  (online)
  //  Bob    → gray dot ○   (offline -- NOT in the Set)
  //  Carol  → green dot ●  (online)

  // ═══════════════════════════════════════════════════════════
  //  TYPING INDICATORS
  // ═══════════════════════════════════════════════════════════
  typingIndicators: {
    "conv_102": ["user_bob"],     // Bob is typing in conv_102
    // conv_101 and conv_103 have no one typing → empty/absent
  },

  // ═══════════════════════════════════════════════════════════
  //  CONNECTION STATUS
  // ═══════════════════════════════════════════════════════════
  connectionStatus: "connected",   // No banner shown (all good)
  // If "connecting" → show "Reconnecting..." yellow banner
  // If "disconnected" → show "No internet connection" red banner

  // ═══════════════════════════════════════════════════════════
  //  PENDING MESSAGES (offline queue)
  // ═══════════════════════════════════════════════════════════
  pendingMessages: [
    // msg_625 is also here because it hasn't been ACKed yet
    {
      id: "msg_625",
      conversationId: "conv_102",
      senderId: "user_me",
      content: "Sure, sending the link?",
      type: "text",
      timestamp: 1710419000000,
      clientTimestamp: 1710419000000,
      status: "sending",
    },
  ],
};
```

### Now see exactly how this mock data renders on screen:

```
┌─────────────────────────────────────────────────────────────────────────┐
│                                                              connected │
│  (no banner shown because connectionStatus = "connected")              │
│                                                                         │
│  ┌────────────────────────────────┐  ┌──────────────────────────────┐  │
│  │         SIDEBAR                │  │       CHAT WINDOW            │  │
│  │                                │  │                              │  │
│  │  [🔍 Search messages...]       │  │  ┌─ ChatHeader ───────────┐ │  │
│  │                                │  │  │  Alice        ● Online │ │  │
│  │  ┌─────────────────────────┐  │  │  │                         │ │  │
│  │  │ ● Alice          5:00PM │  │  │  │  (no typing indicator   │ │  │
│  │  │ Sure, let's meet at 6!  │  │  │  │   because conv_101 not │ │  │
│  │  │              ← selected │  │  │  │   in typingIndicators)  │ │  │
│  │  └─────────────────────────┘  │  │  └─────────────────────────┘ │  │
│  │                                │  │                              │  │
│  │  ┌─────────────────────────┐  │  │  ┌─ Messages ─────────────┐ │  │
│  │  │ ○ Bob         ③ 4:20PM │  │  │  │                         │ │  │
│  │  │ Can you review my PR?   │  │  │  │     Hey Alice!    1:53p │ │  │
│  │  │  ↑ gray dot   ↑ badge  │  │  │  │              ✓✓ blue ──┤ │  │
│  │  │  (offline)    (3 unread)│  │  │  │                    RIGHT│ │  │
│  │  └─────────────────────────┘  │  │  │                         │ │  │
│  │                                │  │  │ Hi! What's up?         │ │  │
│  │  ┌─────────────────────────┐  │  │  │ LEFT ── alice sent it  │ │  │
│  │  │ ● Carol                 │  │  │  │                         │ │  │
│  │  │ (no messages yet)       │  │  │  │  ── Today ──────────── │ │  │
│  │  │  ↑ green dot (online)   │  │  │  │                         │ │  │
│  │  └─────────────────────────┘  │  │  │  Wanna grab dinner     │ │  │
│  │                                │  │  │  tonight?       4:40PM │ │  │
│  │  Sorted by updatedAt:         │  │  │              ✓✓ blue ──┤ │  │
│  │  Alice (5:00) > Bob (4:20)    │  │  │                    RIGHT│ │  │
│  │  > Carol (oldest)             │  │  │                         │ │  │
│  │                                │  │  │ Sure, let's meet       │ │  │
│  │                                │  │  │ at 6!          5:00PM  │ │  │
│  │                                │  │  │ LEFT ── alice sent it  │ │  │
│  │                                │  │  │                         │ │  │
│  │                                │  │  └─────────────────────────┘ │  │
│  │                                │  │                              │  │
│  │                                │  │  ┌─ MessageInput ─────────┐ │  │
│  │                                │  │  │ [+] Type a message...  │ │  │
│  │                                │  │  │                 [Send] │ │  │
│  │                                │  │  └─────────────────────────┘ │  │
│  └────────────────────────────────┘  └──────────────────────────────┘  │
└─────────────────────────────────────────────────────────────────────────┘
```

### If user switches to Bob's chat (activeConversationId = "conv_102"):

```
┌──────────────────────────────────────────────────────────────────────┐
│                                                                      │
│  ┌────────────────────────┐  ┌────────────────────────────────────┐  │
│  │      SIDEBAR           │  │         CHAT WINDOW                │  │
│  │                        │  │                                    │  │
│  │  ● Alice       5:00PM  │  │  ┌─ ChatHeader ─────────────────┐ │  │
│  │  Sure, let's meet...   │  │  │  Bob                ○ Offline │ │  │
│  │                        │  │  │  Bob is typing...             │ │  │
│  │  ○ Bob  ← now selected │  │  │  ↑ from typingIndicators     │ │  │
│  │  Can you review my PR? │  │  │    ["conv_102"] = ["user_bob"]│ │  │
│  │  (badge gone, viewing) │  │  └───────────────────────────────┘ │  │
│  │                        │  │                                    │  │
│  │  ● Carol               │  │  ┌─ Messages ───────────────────┐ │  │
│  │                        │  │  │                               │ │  │
│  └────────────────────────┘  │  │  How's the sprint    3:16 PM │ │  │
│                               │  │  going?          ✓✓ blue ──┤ │  │
│                               │  │                        RIGHT│ │  │
│                               │  │                               │ │  │
│                               │  │  Almost done, just     3:33 │ │  │
│                               │  │  one more ticket             │ │  │
│                               │  │  LEFT ── bob sent it         │ │  │
│                               │  │                               │ │  │
│                               │  │  Nice, let me know    3:50  │ │  │
│                               │  │  if you need help            │ │  │
│                               │  │              ✓✓ gray ───────┤ │  │
│                               │  │  (delivered, not read yet)   │ │  │
│                               │  │                        RIGHT │ │  │
│                               │  │                               │ │  │
│                               │  │  Can you review       4:20  │ │  │
│                               │  │  my PR?                      │ │  │
│                               │  │  LEFT ── bob sent it         │ │  │
│                               │  │                               │ │  │
│                               │  │  Sure, sending the    4:23  │ │  │
│                               │  │  link?                       │ │  │
│                               │  │              ⏳ clock ───────┤ │  │
│                               │  │  (still sending / offline)   │ │  │
│                               │  │                        RIGHT │ │  │
│                               │  │                               │ │  │
│                               │  └───────────────────────────────┘ │  │
│                               │                                    │  │
│                               │  ┌─ MessageInput ───────────────┐ │  │
│                               │  │ [+] Type a message... [Send] │ │  │
│                               │  └───────────────────────────────┘ │  │
│                               └────────────────────────────────────┘  │
└──────────────────────────────────────────────────────────────────────┘

Legend:
  RIGHT bubble = senderId === "user_me"  (my messages, aligned right)
  LEFT  bubble = senderId !== "user_me"  (other person's, aligned left)
  ✓✓ blue  = status: "read"
  ✓✓ gray  = status: "delivered"
  ✓  gray  = status: "sent"
  ⏳ clock = status: "sending"
  ❌ red   = status: "failed" (+ tap to retry)
```

### If connection drops (connectionStatus = "disconnected"):

```
┌──────────────────────────────────────────────────────────────────────┐
│ ⚠️ No internet connection. Messages will be sent when you're back.   │
│ ↑ This red/yellow banner appears because connectionStatus =          │
│   "disconnected"                                                     │
│                                                                      │
│  ┌────────────────────┐  ┌────────────────────────────────────────┐  │
│  │    SIDEBAR          │  │         CHAT WINDOW                    │  │
│  │  (loaded from       │  │  (loaded from IndexedDB cache)         │  │
│  │   IndexedDB cache)  │  │                                        │  │
│  │                     │  │  Messages still visible.                │  │
│  │  User can still     │  │  User can still type and "send."       │  │
│  │  browse and click   │  │  Sent messages show ⏳ (queued).       │  │
│  │  conversations.     │  │  They'll actually send on reconnect.   │  │
│  └─────────────────────┘  └────────────────────────────────────────┘  │
└──────────────────────────────────────────────────────────────────────┘
```

### Quick mapping cheat-sheet:

```
DATA FIELD                          → WHAT THE USER SEES
─────────────────────────────────────────────────────────────────────
conversations["conv_101"]           → One row in the sidebar
  .participants → ["user_alice"]    → Name "Alice" + avatar
  .lastMessage.content              → "Sure, let's meet at 6!" preview
  .unreadCount → 0                  → No badge
  .unreadCount → 3                  → Blue circle with "3"
  .updatedAt                        → Sort order (newest conversation first)

messages["conv_101"]                → All bubbles in the chat window
  [i].senderId === "user_me"        → Bubble on the RIGHT (my message)
  [i].senderId !== "user_me"        → Bubble on the LEFT (their message)
  [i].content                       → Text inside the bubble
  [i].timestamp                     → "5:00 PM" next to the bubble
  [i].status === "read"             → ✓✓ blue checkmarks
  [i].status === "delivered"        → ✓✓ gray checkmarks
  [i].status === "sent"             → ✓ single gray checkmark
  [i].status === "sending"          → ⏳ clock icon
  [i].status === "failed"           → ❌ + "Tap to retry"

onlineUsers.has("user_alice")       → ● Green dot on Alice's avatar
!onlineUsers.has("user_bob")        → ○ Gray dot on Bob's avatar

typingIndicators["conv_102"]        → "Bob is typing..." in chat header
  = ["user_bob"]                      (only visible when viewing conv_102)

connectionStatus === "connected"    → No banner (everything is fine)
connectionStatus === "connecting"   → Yellow banner: "Reconnecting..."
connectionStatus === "disconnected" → Red banner: "No internet connection"

activeConversationId === "conv_101" → Alice's chat is shown in ChatWindow
activeConversationId === null       → "Select a chat to start messaging"

pendingMessages.length > 0          → These get flushed on reconnect
                                      (invisible to user, just ⏳ icons)
```

### Why two timestamps? (Interviewers WILL ask this)

```
Scenario: User hits Send
                          │
                          ▼
           ┌──────────────────────────┐
           │  clientTimestamp = now()  │  ← Set IMMEDIATELY by the client
           │  Show message in UI      │    Used to display the message right away
           │  status = 'sending'      │    (optimistic UI -- don't wait for server)
           └────────────┬─────────────┘
                        │ WebSocket
                        ▼
           ┌──────────────────────────┐
           │  Server receives message │
           │  timestamp = server.now()│  ← Set by the SERVER
           │  Sends ACK back          │    This is the "real" time -- used for
           └────────────┬─────────────┘    ordering across ALL users
                        │
                        ▼
           ┌──────────────────────────┐
           │  Client gets ACK         │
           │  Replace clientTimestamp  │
           │  with server timestamp   │
           │  status = 'sent' ✓       │
           └──────────────────────────┘

Why not just client time?
→ User A's clock is 5 min ahead. User B's is correct.
→ A sends "Hello" at 12:05 (wrong clock), B replies "Hi" at 12:01 (correct clock).
→ If we sort by client time: "Hi" appears BEFORE "Hello". Broken.
→ Server timestamp = single source of truth. Both users see same order.
```

---

## Step 5: Component Design (Minutes 22-27)

### Component Tree (maps to the UI above)

```
<App>
├── <ConnectionBanner />          ← connectionStatus
│
├── <Sidebar>
│   ├── <SearchBar />
│   └── <ConversationList>        ← conversations, onlineUsers
│       └── <ConversationItem />  ← conversation.lastMessage, unreadCount
│
└── <ChatWindow>                  ← activeConversationId
    ├── <ChatHeader />            ← participant name, onlineUsers, typingIndicators
    ├── <MessageList>             ← messages[activeConversationId]
    │   ├── <DateSeparator />     ← "Today", "Yesterday"
    │   ├── <MessageBubble />     ← message.content, status, timestamp, senderId
    │   └── <ScrollToBottom />    ← shown when user scrolled up
    └── <MessageInput>
        ├── <TextArea />          ← auto-growing input
        └── <SendButton />
```

### Key decisions to explain:

> **1. `<MessageBubble>` is stateless.** It receives a `message` prop and renders it. No internal state. This means I can wrap it with `React.memo` -- it only re-renders when the message's `status` changes (sending → sent → delivered → read). Since messages are mostly immutable, most bubbles never re-render.

> **2. `<ConversationList>` and `<MessageList>` use virtualization.** A user might have 500 conversations and 5000 messages. I can't render all of those as DOM nodes. I'd use `@tanstack/virtual` or `react-window` to only render the ~20 items visible on screen.

> **3. `<MessageList>` is the hardest component** -- it handles infinite scroll upward, auto-scroll on new messages, and scroll position preservation. I'll cover this in detail later.

### Cross-question:

> **Interviewer: "How do you handle different message types -- text, image, file?"**

> "I'd use a switch inside the bubble:
>
> ```jsx
> function MessageContent({ message }) {
>   switch (message.type) {
>     case 'text':  return <TextContent text={message.content} />;
>     case 'image': return <ImageContent url={message.content} />;
>     case 'file':  return <FileContent file={message.content} />;
>     default:      return <span>Unsupported message type</span>;
>   }
> }
> ```
>
> The `default` case matters -- if the server sends a type this client version doesn't know, we show a fallback instead of crashing."

---

## Step 6: WebSocket Connection Management (Minutes 27-35)

This is where many candidates are vague. **Go deep here.** This is the backbone of a chat app.

### What to say:

> "The WebSocket connection is the most critical piece. Let me design a robust connection manager."

### WebSocket Manager Design

```typescript
class WebSocketManager {
  private ws: WebSocket | null = null;
  private url: string;
  private reconnectAttempts = 0;
  private maxReconnectAttempts = 10;
  private messageQueue: OutgoingMessage[] = [];
  private heartbeatInterval: number | null = null;

  connect(authToken: string) {
    this.ws = new WebSocket(`${this.url}?token=${authToken}`);

    this.ws.onopen = () => {
      this.reconnectAttempts = 0;    // Reset counter on success
      this.startHeartbeat();          // Keep connection alive
      this.flushQueue();              // Send any queued messages
      store.setConnectionStatus('connected');
    };

    this.ws.onmessage = (event) => {
      const data = JSON.parse(event.data);
      this.handleMessage(data);       // Route to correct handler
    };

    this.ws.onclose = (event) => {
      this.stopHeartbeat();
      store.setConnectionStatus('disconnected');
      if (event.code !== 1000) {      // 1000 = intentional close (logout)
        this.reconnect();             // Unintentional → try to reconnect
      }
    };
  }

  // ── RECONNECTION: exponential backoff + jitter ──────────────
  private reconnect() {
    if (this.reconnectAttempts >= this.maxReconnectAttempts) return;

    //  Attempt 1 → wait 1s
    //  Attempt 2 → wait 2s
    //  Attempt 3 → wait 4s
    //  Attempt 4 → wait 8s  ... up to 30s max
    const delay = Math.min(1000 * Math.pow(2, this.reconnectAttempts), 30000);

    // Add random jitter (±20%) so 10,000 users don't all retry at once
    const jitter = delay * 0.2 * (Math.random() - 0.5);

    this.reconnectAttempts++;
    store.setConnectionStatus('connecting');

    setTimeout(() => this.connect(this.authToken), delay + jitter);
  }

  // ── SENDING: queue if offline ───────────────────────────────
  send(message: OutgoingMessage) {
    if (this.ws?.readyState === WebSocket.OPEN) {
      this.ws.send(JSON.stringify(message));
    } else {
      this.messageQueue.push(message);  // Will be sent on reconnect
    }
  }

  private flushQueue() {
    while (this.messageQueue.length > 0) {
      const msg = this.messageQueue.shift()!;
      this.ws!.send(JSON.stringify(msg));
    }
  }

  // ── HEARTBEAT: detect silent connection death ───────────────
  private startHeartbeat() {
    this.heartbeatInterval = window.setInterval(() => {
      if (this.ws?.readyState === WebSocket.OPEN) {
        this.ws.send(JSON.stringify({ type: 'ping' }));
      }
    }, 30000);  // Every 30 seconds
  }

  // ── ROUTING: incoming events → state updates ────────────────
  private handleMessage(data: ServerMessage) {
    switch (data.type) {
      case 'new_message':      store.addMessage(data.payload);         break;
      case 'message_ack':      store.updateStatus(data.payload);       break;
      case 'delivery_receipt': store.markDelivered(data.payload);      break;
      case 'read_receipt':     store.markRead(data.payload);           break;
      case 'typing':           store.setTyping(data.payload);          break;
      case 'presence':         store.updatePresence(data.payload);     break;
      case 'pong':             /* connection alive, nothing to do */   break;
    }
  }

  disconnect() {
    this.ws?.close(1000, 'User logged out');
  }
}
```

### Explain the 3 key decisions:

#### 1. Exponential Backoff + Jitter

```
Why exponential backoff?
  → If server crashes, don't hammer it with instant retries.
  → Wait longer each time: 1s, 2s, 4s, 8s... gives server time to recover.

Why jitter (random delay)?
  → Imagine server restarts. 10,000 users all disconnect at the same time.
  → Without jitter: all 10,000 retry after exactly 1 second → server crashes again.
  → With jitter: retries spread across 0.8s-1.2s → server handles it fine.
  → This is called the "thundering herd" problem.
```

#### 2. Heartbeat (Ping/Pong)

> "WebSocket connections can **silently die** -- the network drops but `onclose` never fires (common on mobile, WiFi switches). The heartbeat detects this: if we send `ping` and don't get `pong` within 5 seconds, we assume the connection is dead and trigger reconnect."

#### 3. Message Queue

> "When WebSocket is down, `send()` doesn't throw an error -- it silently queues the message. When connection restores, `flushQueue()` sends everything in order. The user never knows the connection briefly dropped."

### Cross-questions:

> **Interviewer: "What happens if the user has the app open in two tabs?"**

> "Each tab opens its own WebSocket. Both receive messages, both update their UI. For notifications, I'd use the **Broadcast Channel API** to coordinate between tabs -- one tab becomes the 'leader' and handles notifications so you don't get duplicates."

> **Interviewer: "How do you handle authentication with WebSocket?"**

> "WebSocket doesn't support custom headers in the browser. Two options:
> 1. **Query param**: `ws://server.com?token=xyz` -- simple but token shows in logs
> 2. **First message auth**: Open connection, send `{type: 'auth', token: 'xyz'}`, server rejects all other messages until auth succeeds
>
> I'd go with option 2 -- more secure."

---

## Step 7: Message Ordering & Delivery Receipts (Minutes 35-42)

### The Message Lifecycle (the MOST important diagram in this whole design)

```
USER A (Sender)                    SERVER                     USER B (Receiver)
     │                                │                              │
     │  1. User presses Send          │                              │
     │  Show in UI immediately        │                              │
     │  status = 'sending' ⏳         │                              │
     │                                │                              │
     │  2. Send via WebSocket ──────► │                              │
     │     {id:'abc', content:'Hi'}   │                              │
     │                                │ 3. Save to DB                │
     │                                │    Assign server timestamp   │
     │                                │                              │
     │  4. ACK ◄───────────────────── │ 5. Forward ──────────────►   │
     │     {id:'abc',                 │    {message: {...}}          │
     │      status:'sent',            │                              │
     │      timestamp: 171040000}     │                              │
     │                                │                              │
     │  Update to 'sent' ✓            │                              │
     │                                │                              │
     │                                │  6. B's client receives it   │
     │                                │     Shows in chat            │
     │                                │     Sends delivery ACK ──►   │
     │                                │                              │
     │  7. ◄────────────────────────  │  ◄───────────────────────    │
     │     delivery_receipt           │                              │
     │     Update to 'delivered' ✓✓   │                              │
     │                                │                              │
     │                                │  8. User B OPENS the chat    │
     │                                │     Sends read ACK ────────► │
     │                                │                              │
     │  9. ◄────────────────────────  │  ◄───────────────────────    │
     │     read_receipt               │                              │
     │     Update to 'read' ✓✓ blue  │                              │
```

### Optimistic UI: Show First, Confirm Later

```typescript
function sendMessage(content: string) {
  const message: Message = {
    id: generateUUID(),              // Client generates the ID
    conversationId: activeConversationId,
    senderId: currentUser.id,
    content,
    type: 'text',
    timestamp: Date.now(),           // Temporary -- replaced by server
    clientTimestamp: Date.now(),
    status: 'sending',               // ⏳
  };

  // 1. Show immediately in UI (don't wait for server)
  store.addMessage(message);

  // 2. Send via WebSocket
  wsManager.send({ type: 'send_message', payload: message });

  // 3. If server doesn't ACK within 10 seconds → mark as failed
  setTimeout(() => {
    if (store.getMessageStatus(message.id) === 'sending') {
      store.updateStatus(message.id, 'failed');   // Show ❌ + Retry button
    }
  }, 10000);
}
```

> "The user sees their message **instantly**. We don't wait for the server round-trip. If the server confirms (ACK), we update the checkmark. If it doesn't confirm within 10 seconds, we show a retry button. This pattern is called **optimistic UI**."

### Message Ordering: Why It's Tricky

> "I use server timestamp as the canonical order. But insertion needs to be efficient:
>
> Most of the time, new messages have the latest timestamp and go at the end of the array (common case, O(1)).
>
> Occasionally, a delayed message arrives out of order. For that, I use binary search to find the correct insertion position (rare case, O(log n))."

### Read Receipts: The Watermark Pattern

> "I don't send a read receipt per message. If User B opens a chat with 50 unread messages, I don't fire 50 events. Instead, I send **one** event: `{lastReadMessageId: 'abc'}`. All messages up to that ID are considered read. This is called the **watermark pattern** -- like a water level that rises."

```typescript
function markConversationAsRead(conversationId: string) {
  const messages = store.getMessages(conversationId);
  const lastMessage = messages[messages.length - 1];

  if (lastMessage && lastMessage.senderId !== currentUser.id) {
    wsManager.send({
      type: 'read_ack',
      payload: { conversationId, lastReadMessageId: lastMessage.id },
    });
  }
}

// Triggered when:
// 1. User clicks on a conversation (switches to it)
// 2. New message arrives AND tab is visible (check with Page Visibility API)
// NOT triggered when tab is in background (don't lie about reading)
```

### Cross-questions:

> **Interviewer: "What if a message fails? How does the user retry?"**

> "Failed messages show ❌ with a 'Tap to retry' button. On retry:
> 1. Reset status to `'sending'`
> 2. Re-send via WebSocket **with the same message ID**
> 3. Why same ID? If the original DID reach the server but the ACK was lost, the server sees the duplicate ID and just re-sends the ACK instead of creating a duplicate message. This is **idempotent sending**."

> **Interviewer: "What about the typing indicator?"**

> "Typing indicators are simple and lossy (by design):
> - **Sender side**: Throttle to 1 event every 3 seconds (not every keystroke)
> - **Receiver side**: Show 'typing...' and auto-clear after 4 seconds of no new event
> - Never persisted anywhere -- purely in-memory, ephemeral"

---

## Step 8: Offline Message Queue (Minutes 42-47)

### What to say:

> "Users expect to type messages on the subway and have them send when they're back online."

### How it works (simple flow):

```
User hits Send while offline
        │
        ▼
  WebSocket is down
  → send() queues message in memory
  → ALSO persist to IndexedDB        ← survives tab close / browser crash
  → UI shows message with ⏳ status
        │
   ... time passes ...
        │
  Connection restored
  → WebSocket reconnects
  → flushQueue(): send all pending messages in order
  → Remove from IndexedDB as each is ACKed by server
  → UI updates ⏳ → ✓
```

### Common Confusion: When does Zustand die? When does IndexedDB kick in?

This is subtle and interviewers may probe on it. Understand this clearly:

**Zustand state lives in RAM (browser memory).** It has NOTHING to do with internet.

```
Going offline (WiFi drops, subway, etc.)
  → Does Zustand state disappear?  NO. It's in RAM. Still there.
  → Does the UI break?             NO. UI reads from Zustand, not from server.
  → Does the page refresh?         NO. Nothing happens to the page.
  → The ONLY thing that happens:
      WebSocket's onclose fires → connectionStatus = "disconnected" → banner shown.
      Everything else on screen stays exactly the same.

Zustand state is ONLY lost when:
  → User closes the tab
  → User refreshes the page (F5)
  → Browser crashes
```

**IndexedDB is the safety net for when Zustand dies (tab close / refresh).**

Think of it as:
```
Zustand    = notepad on your desk    (fast, but gone if you leave the room)
IndexedDB  = photocopy in your drawer (slower, but survives you leaving)
Server     = original in the office   (needs internet to reach)
```

Here are the two scenarios to understand the difference:

```
SCENARIO A: User goes offline, keeps tab open, comes back online
─────────────────────────────────────────────────────────────────
1. Chatting normally                    Zustand: ✅  IndexedDB: ✅ (auto-saved copy)
2. Internet drops                       Zustand: ✅  IndexedDB: ✅
   (UI still works, just shows banner)  (RAM doesn't care about internet)
3. User sends message offline           Zustand: ✅ adds msg ⏳   IndexedDB: ✅ saves copy
4. Internet returns, WebSocket syncs    Zustand: ✅ ⏳ → ✓        IndexedDB: ✅ updates

   IndexedDB was NOT needed here. Tab was never closed. Zustand handled everything.
```

```
SCENARIO B: User goes offline, CLOSES THE TAB, reopens later
─────────────────────────────────────────────────────────────────
1. Chatting normally                    Zustand: ✅  IndexedDB: ✅ (auto-saved copy)
2. Internet drops                       Zustand: ✅  IndexedDB: ✅
3. User sends message offline           Zustand: ✅ adds msg ⏳   IndexedDB: ✅ saves copy
4. User CLOSES THE TAB                  Zustand: ❌ GONE (RAM cleared)
                                        IndexedDB: ✅ still has everything (on disk)
5. User reopens app (still no internet)
   → Zustand starts empty
   → persist middleware reads from IndexedDB → fills Zustand
   → UI shows cached conversations + pending message ⏳
   → No internet needed. IndexedDB is a local file on disk.

   THIS is where IndexedDB saved us. Without it → blank screen.

6. Internet returns eventually
   → WebSocket connects → flushes pending msg → ⏳ becomes ✓
```

**How the auto-sync from Zustand → IndexedDB works (in code):**

```typescript
import { create } from 'zustand';
import { persist } from 'zustand/middleware';

const useChatStore = create(
  persist(                              // ← This middleware handles the sync
    (set) => ({
      conversations: {},
      messages: {},
      addMessage: (msg) => set((state) => ({ /* ...update... */ })),
    }),
    {
      name: 'chat-store',              // key in IndexedDB
      storage: createJSONStorage(() => indexedDBStorage),
    }
  )
);

// What the persist middleware does automatically:
// 1. On EVERY set() call → writes updated state to IndexedDB (background, async)
// 2. On app startup → reads from IndexedDB → fills Zustand with cached data
// You never manually save or load. It just works.
```

**Key interview line to say:**
> "Zustand is the single source of truth at runtime. The persist middleware mirrors state to IndexedDB on every update. IndexedDB is only read on app startup — to rehydrate Zustand after a tab close or refresh. Going offline doesn't affect Zustand at all since it's in-memory."

### Why IndexedDB, not localStorage?

| | localStorage | IndexedDB |
|--|-------------|-----------|
| Sync/Async | **Synchronous** (blocks main thread) | **Async** (non-blocking) |
| Storage limit | 5-10 MB | 50+ MB |
| Data format | Strings only (must JSON.stringify everything) | Structured data natively |
| Queryable | No (key-value only) | Yes (indexes, ranges) |
| **Verdict** | Fine for settings/preferences | **Right tool for chat data** |

### What gets cached offline:

```
IndexedDB: 'chat_app_db'
├── conversations     → Last 20 conversations (metadata)
├── messages          → Last 50 messages per conversation
├── pending_messages  → Unsent messages (the offline queue)
└── user_data         → Current user profile, settings
```

> "When the user opens the app on a slow connection, they immediately see cached conversations and messages from IndexedDB. Then WebSocket connects in the background and syncs new data. The app feels instant."

### What happens on reconnect:

> "Order matters: **receive first, then send**.
> 1. Send a `sync` event with the timestamp of the last message we received
> 2. Server sends back all messages we missed while offline
> 3. Merge those into our local state
> 4. THEN flush our pending outgoing messages
>
> Why this order? So our pending messages get server timestamps that come after the messages we missed. Keeps the ordering correct."

### Cross-question:

> **Interviewer: "What if the user sends messages offline, closes the browser, and opens it next day?"**

> "The pending queue is in IndexedDB, not just memory. On app startup:
> 1. Read `pending_messages` from IndexedDB
> 2. Show them in UI with ⏳ status
> 3. When WebSocket connects, flush them
> 4. Remove from IndexedDB as each is ACKed
>
> The messages survive browser restarts because IndexedDB persists to disk."

---

## Step 9: Scroll Position & Infinite Scroll Up (Minutes 47-53)

### The 3 Scroll Challenges:

```
Challenge 1: Auto-scroll on new message
  → New message arrives while user is at the bottom → scroll down to show it
  → New message arrives while user scrolled UP reading old messages → DON'T scroll
    (show "↓ 3 new messages" button instead)

Challenge 2: Load older messages on scroll up
  → User scrolls to top → fetch older messages from API → prepend them
  → BUT: adding content above current scroll position would push everything down
    → user loses their place → BAD UX
  → Need to RESTORE scroll position after prepending

Challenge 3: Performance
  → Conversation has 10,000 messages → can't render 10,000 DOM nodes
  → Use virtualization → only render ~20-30 visible messages
```

### Solution: Part 1 -- Auto-scroll vs Stay-in-place

```typescript
function MessageList({ conversationId }) {
  const containerRef = useRef<HTMLDivElement>(null);
  const [isNearBottom, setIsNearBottom] = useState(true);
  const messages = useMessages(conversationId);

  // Track: is the user near the bottom?
  const handleScroll = useCallback(() => {
    const el = containerRef.current;
    if (!el) return;
    const distanceFromBottom = el.scrollHeight - el.scrollTop - el.clientHeight;
    setIsNearBottom(distanceFromBottom < 100);  // within 100px = "near bottom"
  }, []);

  // When new message arrives: auto-scroll only if already near bottom
  useEffect(() => {
    if (isNearBottom) {
      containerRef.current?.scrollTo({
        top: containerRef.current.scrollHeight,
        behavior: 'smooth',
      });
    }
  }, [messages.length, isNearBottom]);

  return (
    <div ref={containerRef} onScroll={handleScroll}>
      {messages.map(msg => <MessageBubble key={msg.id} message={msg} />)}

      {/* Floating button when user is scrolled up */}
      {!isNearBottom && (
        <button className="scroll-to-bottom-fab" onClick={scrollToBottom}>
          ↓ 3 new messages
        </button>
      )}
    </div>
  );
}
```

> "The `isNearBottom` flag is the key. WhatsApp, Slack, Telegram -- they all use this exact pattern."

### Solution: Part 2 -- Infinite Scroll Up (the tricky one)

```
Problem visualization:

  BEFORE loading older messages:       AFTER loading older messages:
  ┌────────────────────┐               ┌────────────────────┐
  │ Message 21         │ ← visible     │ Message 1          │ ← NEW (loaded)
  │ Message 22         │ ← visible     │ Message 2          │ ← NEW (loaded)
  │ Message 23         │ ← visible     │ ...                │
  │ Message 24         │ ← visible     │ Message 20         │ ← NEW (loaded)
  │ Message 25         │ ← visible     │ Message 21         │ ← was visible
  └────────────────────┘               │ Message 22         │ ← was visible
                                       │ ...                │
  User is looking at 21-25.            │ Message 25         │
  After prepending, the viewport       └────────────────────┘
  still shows messages 1-5 (top        User is now looking at 1-5!
  of the container). Message 21        Message 21 got pushed way down.
  got pushed down. User lost           
  their place!                         THE FIX: After prepending, set
                                       scrollTop = newScrollHeight - oldScrollHeight
                                       This "pushes" the viewport back down
                                       to where message 21 is.
```

```typescript
async function loadOlderMessages() {
  const el = containerRef.current;
  const previousScrollHeight = el.scrollHeight;         // Save BEFORE

  const olderMessages = await api.getMessages(conversationId, {
    before: messages[0].timestamp,   // "give me messages older than my oldest"
    limit: 30,
  });

  store.prependMessages(conversationId, olderMessages); // Add to state

  // CRITICAL: Restore scroll position
  requestAnimationFrame(() => {
    const newScrollHeight = el.scrollHeight;             // Measure AFTER
    el.scrollTop = newScrollHeight - previousScrollHeight;  // Adjust
  });
}
```

> "The scroll restoration formula: `scrollTop = newScrollHeight - oldScrollHeight`. Before prepending, the container was 5000px tall. After, it's 8000px (3000px of new content above). I set scrollTop to 3000, which puts the viewport exactly where it was before."

### Solution: Part 3 -- Triggering the load

> "I use an **IntersectionObserver** on a hidden sentinel element at the top of the list. When the sentinel enters the viewport (user scrolled to the top), it triggers `loadOlderMessages()`. This is more efficient than listening to scroll events."

### Cross-questions:

> **Interviewer: "What if the user scrolls up AND a new message arrives at the bottom?"**

> "`isNearBottom` is false, so I don't auto-scroll. Instead:
> 1. New message is added to state (appended at the bottom)
> 2. The floating button counter increments: '↓ 3 new messages'
> 3. User clicks it → smooth scroll to bottom, counter resets"

> **Interviewer: "How would you handle 'jump to message' from search results?"**

> "If the message is already in memory, I find its index and scroll to it. If it's not loaded (very old), I fetch messages around that timestamp from the API, load them, then scroll. Either way, I briefly **highlight** the target message with a yellow flash so the user spots it."

---

## Step 10: Wrap-Up & Tradeoffs (Minutes 53-58)

### Tradeoffs Summary (mention this at the end)

| Decision | What I Chose | What I Didn't | Why |
|----------|-------------|---------------|-----|
| Real-time | WebSocket | SSE, long polling | Bidirectional, low overhead |
| State | Zustand | Redux, Context | Lightweight selectors, less boilerplate |
| Offline storage | IndexedDB | localStorage | Async, structured data, 50MB+ |
| Message ordering | Server timestamp | Client timestamp | Consistent across all users |
| Message IDs | Client-generated UUID | Server-assigned | Enables optimistic UI + idempotent retries |
| Virtualization | @tanstack/virtual | react-window | Better dynamic height support |

### "If I had more time, I'd also consider..."

(Say 2-3 of these to show breadth. Don't go deep -- just name-drop.)

- **Accessibility**: `aria-live` region for new messages, keyboard navigation in conversation list
- **Error handling**: Error boundaries around MessageList, graceful fallbacks
- **Media messages**: Upload to S3 → get URL → send URL as message
- **Push notifications**: Service Worker + Push API for when tab is closed
- **End-to-end encryption**: Web Crypto API, but massive scope increase

---

## Cross Questions Bank (with Answers)

### Extending the Design

**Q1: "How would you extend this to group chat?"**
> "Three changes:
> 1. `Conversation` gets `type: 'group' | 'direct'`, participants array has N members
> 2. Read receipts: instead of blue checks, show '3 of 8 read' or individual read avatars
> 3. Typing indicators: 'Alice is typing...' or 'Alice and 2 others...'
>
> WebSocket format stays the same -- server handles fan-out to all group members."

**Q2: "How would you add emoji reactions like Slack?"**
> "Add `reactions: Record<string, string[]>` (emoji → userIds) to Message. Click sends `{type: 'reaction', messageId, emoji}` via WebSocket. Show reaction pills below the bubble."

**Q3: "What if the server goes down?"**
> "WebSocket closes → reconnection attempts with backoff. User sees 'Connecting...' banner. They can still browse cached conversations and send messages (queued). Everything syncs when server is back."

### API & Protocol

**Q4: "REST or GraphQL for the API?"**
> "REST. The data model is simple (conversations, messages). REST gives me HTTP caching. GraphQL's flexibility isn't needed here since queries are predictable."

**Q5: "What if WebSocket is blocked by a firewall?"**
> "Fall back to long polling. Libraries like Socket.IO do this automatically."

**Q6: "What protocol on top of WebSocket?"**
> "JSON for simplicity and debuggability. At scale, Protocol Buffers or MessagePack (~30% smaller) would be better."

### Messages

**Q7: "How do you prevent duplicate messages?"**
> "Client-generated UUIDs. Server uses the ID as an idempotency key -- if it receives the same ID twice (retry), it ignores the duplicate and re-sends the ACK."

**Q8: "How would you do message editing and deletion?"**
> "Edit: server broadcasts `{type: 'message_edited', messageId, newContent}`. UI updates the message and shows '(edited)'.
> Delete: server broadcasts `{type: 'message_deleted', messageId}`. UI replaces content with 'This message was deleted' (don't remove from array, avoids gaps)."

**Q9: "How do you handle images/media?"**
> "1. User selects image → show local preview immediately (URL.createObjectURL)
> 2. Upload to S3 via REST with progress bar
> 3. Once uploaded, send message via WebSocket with the image URL
> 4. Receiver renders thumbnail, click to view full size"

### Performance

**Q10: "500 conversations in the sidebar. How to keep it fast?"**
> "Virtualize the sidebar list too. Only render the ~15-20 visible items. `React.memo` each `ConversationItem` since they rarely change."

**Q11: "User is in 100 active group chats?"**
> "Prioritize: active conversation gets full updates (messages, typing). Others only get lightweight updates (lastMessage, unreadCount). Batch sidebar updates with a 500ms debounce."

### Offline & Sync

**Q12: "How much data to cache offline?"**
> "Bounded: last 20 conversations, last 50 messages each, all pending outgoing messages. Total: ~2-5 MB. LRU eviction if cache grows too large."

**Q13: "User was offline for 3 days?"**
> "Don't dump everything at once. Sync conversation metadata first (lightweight), then fetch messages per-conversation on demand when user opens each chat."

### UX Edge Cases

**Q14: "What about link previews (like Slack)?"**
> "Called 'unfurling'. Server detects URLs, fetches Open Graph metadata (title, description, image), attaches to message. Client renders a preview card. Important: server fetches the URL, not the client (privacy + security)."

**Q15: "User has same chat open on phone and laptop?"**
> "Server sends messages to all active connections for that user. Read receipts from either device sync to the other. Each device syncs independently on reconnect. Eventually consistent."

**Q16: "How would you test this?"**
> "Unit tests: message ordering logic, timestamp formatting (Jest).
> Component tests: MessageBubble for each status (React Testing Library).
> Integration: send/receive flow with mocked WebSocket.
> E2E: two browsers chatting with each other (Playwright)."

---

## Cheat Sheet: Things That Impress Interviewers

### 1. Always explain trade-offs
> Don't just say "I'd use WebSocket." Say "I'd use WebSocket over SSE because chat needs bidirectional communication. The trade-off is more complex server infrastructure."

### 2. Mention failure scenarios unprompted
> "What if the connection drops mid-send? The message is queued. On reconnect, we re-send with the same ID so the server can deduplicate."

### 3. Reference real products
> "WhatsApp uses the watermark pattern for read receipts -- one event marks all messages as read, instead of one event per message."

### 4. Tie decisions back to UX
> "The scroll position restoration exists purely because losing your place while reading is frustrating. It's a small technical detail that makes a huge experience difference."

### 5. Acknowledge what you don't know (this is HUGE)
> "I haven't implemented WebSocket reconnection with jitter in production, but I know the thundering herd problem from reading about it, and here's how I'd approach it..."

This is 10x more impressive than bluffing. The interviewer knows you're at 4 years experience. They want to see **how you think**, not that you've built everything before.

### 6. Mention loading + empty states
> "When user opens a conversation, I'd show skeleton message bubbles while loading. Empty conversation shows 'Say hi to Alice!' as a prompt."

---

## Summary: The 5-Minute Recap

If you only remember 5 things:

1. **WebSocket for real-time, REST for history** -- two channels, each for the right job
2. **Optimistic UI with client-generated IDs** -- show immediately, confirm later, same ID for retries (idempotent)
3. **Exponential backoff + jitter** -- prevents thundering herd on reconnect
4. **Scroll position management** -- auto-scroll if near bottom, preserve position if scrolled up, restore after prepending
5. **IndexedDB for offline queue** -- survives tab close, enables instant app load from cache

---

*Estimated study time: 2-3 hours to read and understand, then 2-3 practice runs explaining it aloud (45 minutes each).*
