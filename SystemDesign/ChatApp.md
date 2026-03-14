# Frontend System Design: Chat Application (Full Interview Guide)

> **Context:** You're in a 45-60 minute frontend system design round. The interviewer says:
> *"Design a chat application like WhatsApp Web or Slack."*
>
> This guide walks you through **exactly what to say, in what order, and why**.
> Written for someone with basic frontend knowledge -- no hand-waving.

---

## Table of Contents

1. [Before We Start: How This Round Actually Works](#before-we-start-how-this-round-actually-works)
2. [Step 1: Requirements Clarification (Minutes 0-5)](#step-1-requirements-clarification-minutes-0-5)
3. [Step 2: High-Level Architecture (Minutes 5-10)](#step-2-high-level-architecture-minutes-5-10)
4. [Step 3: Data Model (Minutes 10-15)](#step-3-data-model-minutes-10-15)
5. [Step 4: Component Design (Minutes 15-22)](#step-4-component-design-minutes-15-22)
6. [Step 5: WebSocket Connection Management (Minutes 22-30)](#step-5-websocket-connection-management-minutes-22-30)
7. [Step 6: Message Ordering & Delivery Receipts (Minutes 30-37)](#step-6-message-ordering--delivery-receipts-minutes-30-37)
8. [Step 7: Offline Message Queue (Minutes 37-42)](#step-7-offline-message-queue-minutes-37-42)
9. [Step 8: Scroll Position & Infinite Scroll Up (Minutes 42-48)](#step-8-scroll-position--infinite-scroll-up-minutes-42-48)
10. [Step 9: Polish & Tradeoffs (Minutes 48-55)](#step-9-polish--tradeoffs-minutes-48-55)
11. [Cross Questions Bank (with Answers)](#cross-questions-bank-with-answers)
12. [Cheat Sheet: Things That Impress Interviewers](#cheat-sheet-things-that-impress-interviewers)

---

## Before We Start: How This Round Actually Works

### What the interviewer is REALLY evaluating:

| Signal | What They Look For | Weight |
|--------|-------------------|--------|
| **Structured thinking** | Do you have a framework? Do you go step-by-step? | Very High |
| **Requirement gathering** | Do you ask questions or just start designing? | High |
| **Trade-off awareness** | Can you explain WHY you chose X over Y? | Very High |
| **Depth in key areas** | Can you go deep when probed? | High |
| **Communication** | Do you think aloud? Can you draw/explain clearly? | High |
| **Real-world awareness** | Do you mention edge cases, failures, scale? | Medium |

### What they are NOT evaluating:
- Memorized system designs
- Perfect answers
- Knowing every API by name

### Golden rule:
**Talk like you're designing this with a teammate, not reciting an answer to a teacher.**

Say things like:
- "One approach would be... but the trade-off is..."
- "I'd lean towards X because in my experience..."
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
| "Is there message editing/deletion?" | Affects data model | "Nice to have" |

### Non-Functional Requirements (state these yourself):

> "For non-functional requirements, I'll assume:
> - **Low latency**: messages should appear in < 200ms under normal conditions
> - **Reliability**: no messages should be lost, even on flaky networks
> - **Message ordering**: messages must appear in the correct order
> - **Offline resilience**: users should be able to queue messages when offline
> - **Scalable**: should handle a conversation with thousands of messages (scroll performance)
> - **Accessible**: keyboard navigable, screen reader friendly"

### After clarification, summarize:

> "So to summarize, I'm designing a 1-on-1 chat app with real-time messaging, delivery/read receipts, typing indicators, offline support, and text messages. I'll mention how to extend to group chat and media. Let me start with the high-level architecture."

**Why this matters:** You just showed structured thinking and communication. Many candidates get rejected because they spend 40 minutes designing the wrong thing.

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
│  │              State Management Layer               │ │
│  │         (React Context / Zustand / Redux)         │ │
│  └──────────────────────┬────────────────────────────┘ │
│                         │                              │
│  ┌──────────────────────▼────────────────────────────┐ │
│  │              Local Storage / IndexedDB             │ │
│  │            (offline cache, message queue)          │ │
│  └───────────────────────────────────────────────────┘ │
└──────────────────────────┬───────────────────────────┘
                           │
              ┌────────────▼────────────┐
              │     WebSocket Server    │
              │   (persistent connection)│
              └────────────┬────────────┘
              ┌────────────▼────────────┐
              │      REST API Server    │
              │  (auth, history, media) │
              └────────────┬────────────┘
              ┌────────────▼────────────┐
              │        Database         │
              └─────────────────────────┘
```

### What to say:

> "I see three key communication channels between client and server:
>
> 1. **WebSocket connection** -- for real-time message delivery, typing indicators, presence/online status, and read receipts. This is a persistent, bi-directional connection.
>
> 2. **REST API** -- for non-real-time operations: authentication, fetching message history (pagination), user profile, uploading media. We don't need WebSocket for these because they're request-response.
>
> 3. **Service Worker** -- for offline support. It intercepts failed requests, queues messages in IndexedDB, and retries when connection is restored.
>
> On the client side, I have a state management layer that acts as the single source of truth for all chat data. The UI subscribes to this state."

### Cross-question you might get here:

> **Interviewer: "Why WebSocket and not something else like Server-Sent Events or long polling?"**

Your answer:

> "Good question. Let me compare the three options:
>
> - **Long Polling**: Client keeps making HTTP requests. Works everywhere but is inefficient -- each poll is a new TCP connection, high latency, wastes bandwidth. Acceptable as a fallback.
>
> - **Server-Sent Events (SSE)**: Server pushes events to the client over a single HTTP connection. It's simpler than WebSocket but **unidirectional** -- the client can't send messages through the SSE channel. For chat, we need bidirectional communication, so SSE alone isn't enough. We'd still need a separate channel for sending messages.
>
> - **WebSocket**: Persistent, bidirectional, low overhead after the initial handshake. Perfect for chat where both sides are constantly sending data.
>
> I'd go with **WebSocket as primary**, with **long polling as a fallback** for environments where WebSocket is blocked (some corporate firewalls)."

---

## Step 3: Data Model (Minutes 10-15)

### What to say:

> "Let me define the data shapes we'll be working with on the client side."

### Message Object

```typescript
interface Message {
  id: string;               // Unique ID (UUID generated on client)
  conversationId: string;   // Which chat thread
  senderId: string;         // Who sent it
  content: string;          // Message body
  type: 'text' | 'image' | 'file';  // For future extensibility
  timestamp: number;        // Unix ms -- server-assigned for ordering
  clientTimestamp: number;   // Client-local time when composed
  status: 'sending' | 'sent' | 'delivered' | 'read' | 'failed';
  replyTo?: string;         // ID of message being replied to (optional)
}
```

### Why TWO timestamps?

> "I'm keeping both `clientTimestamp` and `timestamp` deliberately:
>
> - `clientTimestamp` is set when the user hits send. It's used for **optimistic UI** -- I show the message immediately in the chat, sorted by this time, so the user gets instant feedback.
>
> - `timestamp` is assigned by the **server** when it receives the message. This is the source of truth for **ordering across users**. Once the server responds, I update the message with the server timestamp.
>
> Why not just use client time? Because **client clocks can be wrong**. If User A's clock is 5 minutes ahead, their messages would appear out of order for User B. The server acts as the single source of truth for ordering."

### Conversation Object

```typescript
interface Conversation {
  id: string;
  participants: string[];        // User IDs
  lastMessage: Message | null;   // For conversation list preview
  unreadCount: number;
  updatedAt: number;             // For sorting conversations
}
```

### Client-Side State Shape

```typescript
interface ChatState {
  // Current user
  currentUser: User;

  // All conversations, keyed by ID for O(1) lookup
  conversations: Record<string, Conversation>;

  // Messages per conversation, keyed by conversation ID
  // Each is an array sorted by timestamp
  messages: Record<string, Message[]>;

  // Active conversation
  activeConversationId: string | null;

  // Online status of users
  onlineUsers: Set<string>;

  // Who is currently typing in which conversation
  typingIndicators: Record<string, string[]>; // convId -> [userId, ...]

  // Network state
  connectionStatus: 'connected' | 'connecting' | 'disconnected';

  // Pending messages waiting to be sent (offline queue)
  pendingMessages: Message[];
}
```

### What to explain about this shape:

> "A few intentional choices here:
>
> 1. **Messages are normalized by conversation ID** -- I don't flatten all messages into one giant array. This means switching conversations is O(1) lookup, and I only load messages for conversations the user opens.
>
> 2. **`conversations` is a Record (hash map), not an array** -- because I'll frequently need to update a specific conversation when a new message arrives. With an array, that's O(n) to find. With a Record, it's O(1).
>
> 3. **`onlineUsers` is a Set** -- O(1) lookup for checking if someone is online.
>
> 4. **`pendingMessages` is separate** -- these are messages the user sent while offline. They need special treatment (retry logic, eventual reconciliation with server)."

### Cross-question:

> **Interviewer: "Would you use Redux, Zustand, or Context for this state?"**

Your answer:

> "For a chat app, I'd avoid plain Context API because **any state change in context re-renders all subscribers**. Chat state changes very frequently (every incoming message), so this would cause performance issues.
>
> I'd lean towards **Zustand** or **Redux Toolkit**:
> - Zustand is simpler, has built-in selector support so components only re-render when their specific slice changes. Good for most cases.
> - Redux Toolkit if the team is already familiar with Redux, or if we need powerful devtools and middleware (logging, persistence).
>
> For this design, I'll go with Zustand for simplicity, with middleware for IndexedDB persistence."

---

## Step 4: Component Design (Minutes 15-22)

### Component Tree

```
<App>
├── <Sidebar>
│   ├── <UserProfile />           -- Current user avatar, status
│   ├── <SearchBar />             -- Search conversations
│   └── <ConversationList>
│       └── <ConversationItem />  -- Avatar, name, last message preview, unread badge
│
├── <ChatWindow>
│   ├── <ChatHeader />            -- Other user's name, online status, typing indicator
│   ├── <MessageList>             -- The scrollable message area
│   │   ├── <DateSeparator />     -- "Today", "Yesterday", "March 14"
│   │   ├── <MessageBubble />     -- Individual message
│   │   │   ├── <MessageContent />
│   │   │   ├── <MessageStatus /> -- ✓ sent, ✓✓ delivered, ✓✓ (blue) read
│   │   │   └── <MessageTime />
│   │   └── <ScrollToBottom />    -- FAB button when scrolled up
│   └── <MessageInput>
│       ├── <TextArea />          -- Auto-growing input
│       ├── <EmojiPicker />
│       └── <SendButton />
│
└── <ConnectionBanner />          -- "Reconnecting..." when offline
```

### Key Component Decisions to Explain:

> "A few important design decisions in the component structure:
>
> **1. `<MessageList>` is the most complex component.** It handles:
> - Rendering potentially thousands of messages efficiently (virtualization)
> - Infinite scroll upward to load older messages
> - Maintaining scroll position when new messages are loaded above
> - Auto-scrolling to bottom when a new message arrives (but only if the user is already at the bottom)
>
> **2. `<MessageBubble>` is intentionally simple.** It receives a `message` object and renders it. It doesn't manage any state. This makes it easy to memoize with `React.memo` -- since messages are immutable (once received), the component won't re-render unless the message status changes (sending → sent → delivered → read).
>
> **3. `<ConversationList>` and `<MessageList>` should both use virtualization** because a user might have hundreds of conversations and thousands of messages. I'd use something like `react-window` or `@tanstack/virtual`."

### Cross-question:

> **Interviewer: "How would you handle the MessageBubble for different message types -- text, image, file, reply?"**

Your answer:

> "I'd use a **composition pattern** with a content renderer:
>
> ```jsx
> function MessageBubble({ message }) {
>   return (
>     <div className={message.senderId === currentUser.id ? 'sent' : 'received'}>
>       {message.replyTo && <ReplyPreview messageId={message.replyTo} />}
>       <MessageContent message={message} />
>       <MessageMeta time={message.timestamp} status={message.status} />
>     </div>
>   );
> }
>
> function MessageContent({ message }) {
>   switch (message.type) {
>     case 'text':  return <TextContent text={message.content} />;
>     case 'image': return <ImageContent url={message.content} />;
>     case 'file':  return <FileContent file={message.content} />;
>     default:      return <UnsupportedContent />;
>   }
> }
> ```
>
> The `default` case is important -- if the server sends a message type this client version doesn't support (say 'voice'), we gracefully show 'Unsupported message type' rather than crashing."

---

## Step 5: WebSocket Connection Management (Minutes 22-30)

This is where many candidates are vague. **Go deep here.** This is the backbone of a chat app.

### What to say:

> "The WebSocket connection is the most critical piece of infrastructure in this app. Let me design a robust connection manager."

### WebSocket Manager Design

```typescript
class WebSocketManager {
  private ws: WebSocket | null = null;
  private url: string;
  private reconnectAttempts = 0;
  private maxReconnectAttempts = 10;
  private messageQueue: OutgoingMessage[] = []; // Buffer during disconnection
  private heartbeatInterval: number | null = null;
  private listeners: Map<string, Set<Function>> = new Map();

  constructor(url: string) {
    this.url = url;
  }

  connect(authToken: string) {
    // Pass auth token as query param or in first message
    this.ws = new WebSocket(`${this.url}?token=${authToken}`);

    this.ws.onopen = () => {
      this.reconnectAttempts = 0;            // Reset on successful connect
      this.startHeartbeat();                  // Keep connection alive
      this.flushQueue();                      // Send any queued messages
      this.emit('connectionChange', 'connected');
    };

    this.ws.onmessage = (event) => {
      const data = JSON.parse(event.data);
      this.handleMessage(data);
    };

    this.ws.onclose = (event) => {
      this.stopHeartbeat();
      this.emit('connectionChange', 'disconnected');

      // Don't reconnect if closed intentionally (code 1000)
      if (event.code !== 1000) {
        this.reconnect();
      }
    };

    this.ws.onerror = () => {
      // onerror is always followed by onclose, so reconnect logic lives there
    };
  }

  private reconnect() {
    if (this.reconnectAttempts >= this.maxReconnectAttempts) {
      this.emit('connectionFailed', 'Max reconnection attempts reached');
      return;
    }

    // Exponential backoff: 1s, 2s, 4s, 8s, 16s... capped at 30s
    const delay = Math.min(1000 * Math.pow(2, this.reconnectAttempts), 30000);

    // Add jitter (±20%) to prevent thundering herd
    const jitter = delay * 0.2 * (Math.random() - 0.5);

    this.reconnectAttempts++;
    this.emit('connectionChange', 'connecting');

    setTimeout(() => this.connect(this.authToken), delay + jitter);
  }

  send(message: OutgoingMessage) {
    if (this.ws?.readyState === WebSocket.OPEN) {
      this.ws.send(JSON.stringify(message));
    } else {
      // Connection is down -- queue the message
      this.messageQueue.push(message);
    }
  }

  private flushQueue() {
    while (this.messageQueue.length > 0) {
      const msg = this.messageQueue.shift()!;
      this.ws!.send(JSON.stringify(msg));
    }
  }

  // Server expects a ping every 30s, otherwise it kills the connection
  private startHeartbeat() {
    this.heartbeatInterval = window.setInterval(() => {
      if (this.ws?.readyState === WebSocket.OPEN) {
        this.ws.send(JSON.stringify({ type: 'ping' }));
      }
    }, 30000);
  }

  private stopHeartbeat() {
    if (this.heartbeatInterval) {
      clearInterval(this.heartbeatInterval);
    }
  }

  private handleMessage(data: ServerMessage) {
    switch (data.type) {
      case 'new_message':
        this.emit('newMessage', data.payload);
        break;
      case 'message_ack':
        // Server confirmed receipt -- update status to 'sent'
        this.emit('messageAck', data.payload);
        break;
      case 'delivery_receipt':
        this.emit('deliveryReceipt', data.payload);
        break;
      case 'read_receipt':
        this.emit('readReceipt', data.payload);
        break;
      case 'typing':
        this.emit('typingIndicator', data.payload);
        break;
      case 'presence':
        this.emit('presenceUpdate', data.payload);
        break;
      case 'pong':
        // Heartbeat response -- connection is alive
        break;
    }
  }

  disconnect() {
    this.ws?.close(1000, 'User logged out');   // 1000 = normal closure
  }
}
```

### Walk the interviewer through the KEY decisions:

#### 1. Reconnection with Exponential Backoff + Jitter

> "When the connection drops, I don't retry immediately or at fixed intervals. I use **exponential backoff** -- wait 1s, then 2s, then 4s, then 8s, up to 30s max.
>
> I also add **jitter** (random variation). Why? Imagine 10,000 users lose connection at the same time (server restart). Without jitter, all 10,000 would try to reconnect after exactly 1 second -- a thundering herd that could crash the server again. Jitter spreads them out randomly."

#### 2. Heartbeat / Keep-Alive

> "WebSocket connections can silently die -- the network can drop without triggering `onclose` (especially on mobile). To detect this, I send a `ping` every 30 seconds. If the server doesn't respond with `pong` within a timeout, I consider the connection dead and trigger reconnect.
>
> This also keeps the connection alive through proxies and load balancers that might kill idle connections."

#### 3. Message Queue During Disconnection

> "When the WebSocket is down, calls to `send()` don't throw an error -- they queue the message. When the connection is restored, `flushQueue()` sends all pending messages in order. This gives the user a seamless experience -- they can keep typing and sending even during brief disconnections."

#### 4. Connection State in UI

> "I expose the connection state (`connected`, `connecting`, `disconnected`) to the UI. When disconnected, I show a banner: 'Reconnecting...' with a subtle animation. This manages user expectations."

### Cross-questions:

> **Interviewer: "What happens if the user has the app open in two tabs?"**

> "Each tab would open its own WebSocket connection. This is fine from the server's perspective -- it just needs to deliver messages to all active connections for that user.
>
> But on the client, I need to be careful about **notifications**. I don't want both tabs showing a notification for the same message. Options:
> 1. Use the **Broadcast Channel API** to coordinate between tabs. One tab becomes the 'leader' and handles notifications.
> 2. Or, use a **Shared Worker** to maintain a single WebSocket connection shared across all tabs.
>
> For simplicity, I'd start with option 1. The leader tab shows notifications; other tabs just update their UI."

> **Interviewer: "How do you handle authentication with WebSocket?"**

> "WebSocket doesn't support custom headers in the initial handshake (browser limitation). So I have two options:
>
> 1. **Pass the token as a query parameter**: `ws://server.com?token=xyz`. Simple but the token appears in server logs and URL. Use short-lived tokens.
> 2. **Authenticate in the first message**: Open the connection, then send an `{type: 'auth', token: 'xyz'}` message. The server doesn't accept other messages until auth succeeds.
>
> Option 2 is more secure. I'd use that."

---

## Step 6: Message Ordering & Delivery Receipts (Minutes 30-37)

### The Message Lifecycle

> "Let me walk through the complete lifecycle of a message, from send to read."

```
USER A (Sender)                          SERVER                          USER B (Receiver)
     │                                      │                                  │
     │  1. User presses Send                │                                  │
     │  ── Show message with status ──►     │                                  │
     │     'sending' (gray clock icon)      │                                  │
     │                                      │                                  │
     │  2. Send via WebSocket ─────────►    │                                  │
     │     {type:'send_message',            │                                  │
     │      id:'abc', content:'Hi'}         │                                  │
     │                                      │  3. Server persists message      │
     │                                      │     Assigns server timestamp     │
     │                                      │                                  │
     │  4. Server ACKs ◄───────────────     │  5. Server forwards ──────────►  │
     │     {type:'message_ack',             │     {type:'new_message',         │
     │      id:'abc', status:'sent',        │      message: {...}}             │
     │      timestamp: 1710400000}          │                                  │
     │                                      │                                  │
     │  ── Update to 'sent' ──►             │                                  │
     │     (single gray check ✓)            │                                  │
     │                                      │                                  │
     │                                      │  6. Client B receives message    │
     │                                      │     Shows in chat, sends ACK ──► │
     │                                      │     {type:'delivery_ack',        │
     │                                      │      messageId:'abc'}            │
     │                                      │                                  │
     │  7. Delivery receipt ◄──────────     │  ◄───────────────────────────     │
     │     {type:'delivery_receipt',        │                                  │
     │      messageId:'abc'}                │                                  │
     │                                      │                                  │
     │  ── Update to 'delivered' ──►        │                                  │
     │     (double gray check ✓✓)           │                                  │
     │                                      │                                  │
     │                                      │  8. User B opens/views the       │
     │                                      │     conversation, sends ─────►   │
     │                                      │     {type:'read_ack',            │
     │                                      │      conversationId:'xyz',       │
     │                                      │      lastReadMessageId:'abc'}    │
     │                                      │                                  │
     │  9. Read receipt ◄──────────────     │  ◄───────────────────────────     │
     │     {type:'read_receipt',            │                                  │
     │      conversationId:'xyz',           │                                  │
     │      lastReadMessageId:'abc'}        │                                  │
     │                                      │                                  │
     │  ── Update to 'read' ──►             │                                  │
     │     (double blue check ✓✓)           │                                  │
```

### Message Ordering: The Hard Problem

> "Message ordering seems simple but has subtle edge cases."

#### Problem: Why can't we just sort by timestamp?

> "Consider this scenario:
> - User A sends 'Hello' at 12:00:01
> - User B sends 'Hi' at 12:00:02
> - Due to network latency, the server receives B's message first
>
> If I sort by **client timestamp**, the order is correct on each client but might differ between them (clock skew). If I sort by **server timestamp**, both clients see the same order, but it might not match what users expect.
>
> **My approach:** Use server timestamp as the **canonical order**. But within a single sender's messages, also preserve **client-side order** (messages from the same sender should never appear out of sequence relative to each other, since the sender knows their own order)."

#### Implementation:

```typescript
function insertMessageInOrder(messages: Message[], newMessage: Message): Message[] {
  // Binary search for insert position based on server timestamp
  let low = 0, high = messages.length;

  while (low < high) {
    const mid = Math.floor((low + high) / 2);
    if (messages[mid].timestamp <= newMessage.timestamp) {
      low = mid + 1;
    } else {
      high = mid;
    }
  }

  // Insert at the correct position
  const result = [...messages];
  result.splice(low, 0, newMessage);
  return result;
}
```

> "I use **binary search** for insertion rather than appending and re-sorting, because the common case is that new messages have the latest timestamp (so they go at the end in O(1)), and out-of-order messages are rare. Binary search is O(log n) for the rare case."

### Optimistic UI for Sending

```typescript
function sendMessage(content: string) {
  const optimisticMessage: Message = {
    id: generateUUID(),           // Client-generated
    conversationId: activeConversation.id,
    senderId: currentUser.id,
    content,
    type: 'text',
    timestamp: Date.now(),         // Temporary, will be replaced by server
    clientTimestamp: Date.now(),
    status: 'sending',
  };

  // 1. Immediately add to UI (optimistic)
  addMessage(optimisticMessage);

  // 2. Send via WebSocket
  wsManager.send({
    type: 'send_message',
    payload: optimisticMessage,
  });

  // 3. Set a timeout for failure detection
  setTimeout(() => {
    const msg = getMessageById(optimisticMessage.id);
    if (msg?.status === 'sending') {
      // Server didn't ACK in time -- mark as failed
      updateMessageStatus(optimisticMessage.id, 'failed');
    }
  }, 10000); // 10 second timeout
}

// When server ACKs:
function handleMessageAck(ack: { id: string; timestamp: number }) {
  updateMessage(ack.id, {
    timestamp: ack.timestamp,    // Replace with server timestamp
    status: 'sent',
  });
}
```

### Read Receipts: Efficient Implementation

> "I don't send a read receipt for every individual message. That would be noisy -- if User B opens a conversation with 50 unread messages, I don't want 50 WebSocket events.
>
> Instead, I send a **single read receipt with the last read message ID**. The server then marks all messages up to that ID as read. This is a **watermark pattern**."

```typescript
function markConversationAsRead(conversationId: string) {
  const messages = getMessages(conversationId);
  const lastMessage = messages[messages.length - 1];

  if (lastMessage && lastMessage.senderId !== currentUser.id) {
    wsManager.send({
      type: 'read_ack',
      payload: {
        conversationId,
        lastReadMessageId: lastMessage.id,
      },
    });
  }
}
```

> "I trigger this when:
> - The user switches to a conversation (clicks on it)
> - The user is already in a conversation and a new message arrives AND the tab is visible (use the **Page Visibility API** to check -- don't mark as read if the tab is in background)"

### Cross-questions:

> **Interviewer: "What if a message fails to send? How does the user retry?"**

> "Failed messages show a red exclamation icon with a 'Retry' button. On tap:
>
> 1. Reset status to `'sending'`
> 2. Re-send via WebSocket (same message ID, so the server can deduplicate)
> 3. Reset the timeout
>
> The message keeps its original position in the chat (based on `clientTimestamp`) so it doesn't jump around. Using the **same message ID** is key -- if the original actually did reach the server but the ACK was lost, the server sees the duplicate ID and just re-sends the ACK instead of creating a duplicate message. This is called **idempotent message sending**."

> **Interviewer: "How do you handle the typing indicator?"**

> "Typing indicators are fire-and-forget -- they don't need delivery guarantees. I send a `{type: 'typing', conversationId: 'xyz'}` event via WebSocket when the user types.
>
> Key details:
> - **Throttle** the typing event to at most once every 2-3 seconds (don't send on every keystroke)
> - The receiving client shows 'typing...' and sets a **timeout of ~4 seconds**. If no new typing event arrives, clear the indicator. This handles the case where the user stopped typing but we didn't get a 'stopped_typing' event (unreliable network).
> - Don't persist typing indicators to any database -- they're purely ephemeral, in-memory state."

---

## Step 7: Offline Message Queue (Minutes 37-42)

### What to say:

> "For a chat app, offline support isn't optional -- users expect to be able to write messages on the subway and have them send when they're back online. Let me explain my approach."

### Architecture

```
User sends message while offline
        │
        ▼
┌─────────────────────────────┐
│  WebSocket is disconnected  │
│  send() queues message      │
│  in wsManager.messageQueue  │
└─────────────┬───────────────┘
              │
              ▼
┌─────────────────────────────┐
│  Also persist to IndexedDB  │  ◄── Survives tab close / browser crash
│  (pending_messages store)   │
└─────────────┬───────────────┘
              │
              ▼
┌─────────────────────────────┐
│  UI shows message with      │
│  status: 'sending'          │
│  (gray clock icon)          │
└─────────────────────────────┘

... time passes, connection is restored ...

┌─────────────────────────────┐
│  WebSocket reconnects       │
│  onopen fires               │
└─────────────┬───────────────┘
              │
              ▼
┌─────────────────────────────┐
│  flushQueue():              │
│  - Read pending from        │
│    IndexedDB                │
│  - Send each in order       │
│  - Remove from IndexedDB    │
│    as each is ACKed         │
└─────────────────────────────┘
```

### Why IndexedDB and Not localStorage?

> "**localStorage** is synchronous and blocks the main thread, has a 5-10 MB limit, and can only store strings. For a chat app with potentially many queued messages and cached conversation data, those are dealbreakers.
>
> **IndexedDB** is asynchronous, supports structured data (no JSON.stringify needed), has much larger storage limits (~50MB+ depending on browser), and supports indexes for efficient querying. It's the right tool for offline data in a chat app."

### What Gets Cached Offline

```
IndexedDB Database: 'chat_app'
├── Store: 'messages'          -- Recent messages per conversation (last 50)
│   Index: by conversationId
│   Index: by timestamp
├── Store: 'conversations'     -- All conversation metadata
├── Store: 'pending_messages'  -- Messages queued for sending
│   Index: by conversationId
│   Index: by clientTimestamp
└── Store: 'user_data'         -- Current user profile, settings
```

> "When the user opens the app offline (or before the WebSocket connects), they see:
> - Their conversation list (from cached `conversations`)
> - The last 50 messages of each conversation (from cached `messages`)
> - Any pending messages they'd sent earlier (from `pending_messages`)
>
> This means the app loads instantly, even on slow connections -- we show cached data first, then sync in the background."

### Handling Conflicts on Reconnect

> "When we come back online, there might be messages that arrived while we were offline (sent by other users to us). The server would have queued these. On reconnect:
>
> 1. We send a `{type: 'sync', lastReceivedTimestamp: 1710400000}` event
> 2. The server sends all messages we missed since that timestamp
> 3. We merge them into our local state and IndexedDB cache
> 4. We also flush our pending outgoing messages
>
> The order of operations matters: **receive first, then send**. This way, our pending messages get the correct server timestamps relative to messages we missed."

### Cross-question:

> **Interviewer: "What if the user sends a message offline, closes the browser, and reopens it later?"**

> "This is exactly why I persist the pending queue to IndexedDB, not just in-memory. On app startup:
>
> 1. Read `pending_messages` from IndexedDB
> 2. Show them in the UI with `status: 'sending'`
> 3. When WebSocket connects, flush them to the server
> 4. Remove from IndexedDB as they're ACKed
>
> If the app was a PWA with a Service Worker, the Service Worker could even attempt to send messages via the **Background Sync API** -- the browser would retry sending even if the user has closed the tab. But that's a progressive enhancement, not a requirement."

---

## Step 8: Scroll Position & Infinite Scroll Up (Minutes 42-48)

### The Problem

> "Message lists have one of the trickiest scrolling behaviors in all of frontend development. Let me break down the challenges:
>
> 1. **New messages should auto-scroll to bottom** -- but only if the user is already at the bottom
> 2. **If the user scrolled up to read older messages**, a new message should NOT yank them down
> 3. **Loading older messages (scroll up)** should NOT change the visible scroll position -- the user should stay looking at the same message
> 4. **The list could have thousands of messages** -- we need virtualization"

### Solution: Break it Into Parts

#### Part 1: Auto-scroll vs Stay-in-place

```typescript
function MessageList({ conversationId }) {
  const containerRef = useRef<HTMLDivElement>(null);
  const [isNearBottom, setIsNearBottom] = useState(true);
  const messages = useMessages(conversationId);

  // Track if user is near bottom
  const handleScroll = useCallback(() => {
    const el = containerRef.current;
    if (!el) return;

    const threshold = 100; // px from bottom
    const distanceFromBottom = el.scrollHeight - el.scrollTop - el.clientHeight;
    setIsNearBottom(distanceFromBottom < threshold);
  }, []);

  // Auto-scroll to bottom when new message arrives (only if already near bottom)
  useEffect(() => {
    if (isNearBottom) {
      containerRef.current?.scrollTo({
        top: containerRef.current.scrollHeight,
        behavior: 'smooth',
      });
    }
  }, [messages.length, isNearBottom]);

  return (
    <div ref={containerRef} onScroll={handleScroll} className="message-list">
      {messages.map(msg => <MessageBubble key={msg.id} message={msg} />)}
      {!isNearBottom && <ScrollToBottomFAB onClick={scrollToBottom} />}
    </div>
  );
}
```

> "The key insight is the `isNearBottom` flag. I check if the user is within 100px of the bottom. If yes, new messages auto-scroll. If no, I show a floating 'scroll to bottom' button with an unread count badge -- just like WhatsApp."

#### Part 2: Infinite Scroll Upward (Load Older Messages)

> "This is the tricky part. When the user scrolls to the top, I fetch older messages. But inserting content above the current scroll position would push everything down, making the user lose their place."

```typescript
function useInfiniteScrollUp(containerRef, conversationId) {
  const [isLoadingOlder, setIsLoadingOlder] = useState(false);
  const [hasMore, setHasMore] = useState(true);

  const loadOlderMessages = useCallback(async () => {
    if (isLoadingOlder || !hasMore) return;

    setIsLoadingOlder(true);

    const el = containerRef.current;
    const previousScrollHeight = el.scrollHeight;

    // Fetch older messages (before the oldest message we have)
    const oldestMessage = messages[0];
    const olderMessages = await api.getMessages(conversationId, {
      before: oldestMessage.timestamp,
      limit: 30,
    });

    if (olderMessages.length < 30) {
      setHasMore(false); // No more history
    }

    // Prepend to state
    prependMessages(conversationId, olderMessages);

    // CRITICAL: Restore scroll position after DOM update
    // New content was added above, so scrollHeight increased.
    // To keep the same messages visible, adjust scrollTop by the difference.
    requestAnimationFrame(() => {
      const newScrollHeight = el.scrollHeight;
      el.scrollTop = newScrollHeight - previousScrollHeight;
    });

    setIsLoadingOlder(false);
  }, [conversationId, isLoadingOlder, hasMore, messages]);

  // Trigger load when scrolled to top
  useEffect(() => {
    const el = containerRef.current;
    if (!el) return;

    const observer = new IntersectionObserver(
      ([entry]) => {
        if (entry.isIntersecting) {
          loadOlderMessages();
        }
      },
      { root: el, threshold: 0.1 }
    );

    // Observe a sentinel element at the top of the list
    const sentinel = el.querySelector('.scroll-sentinel-top');
    if (sentinel) observer.observe(sentinel);

    return () => observer.disconnect();
  }, [loadOlderMessages]);
}
```

> "The critical line is the scroll position restoration:
>
> ```
> el.scrollTop = newScrollHeight - previousScrollHeight;
> ```
>
> Before loading, I save `scrollHeight`. After new messages are prepended, `scrollHeight` increases. I set `scrollTop` to the difference, which keeps the user looking at exactly the same message they were looking at. I do this in a `requestAnimationFrame` to ensure the DOM has updated."

#### Part 3: Virtualization for Performance

> "If a conversation has 10,000 messages, we can't render all 10,000 DOM nodes. The browser would freeze. I'd use **virtualization** -- only render the ~20-30 messages visible in the viewport, plus a small buffer above and below.
>
> For this, I'd use a library like `react-window` or `@tanstack/virtual`. The tricky part with virtualization in a chat app is that **message heights vary** (some are one line, some are paragraphs, some have images). So I'd use a **dynamic-height virtualizer** that measures items after render.
>
> `@tanstack/virtual` handles this well with its `measureElement` callback."

### Cross-questions:

> **Interviewer: "What if the user scrolls up, loads older messages, and a new message arrives at the bottom?"**

> "Since `isNearBottom` is false (user scrolled up), I don't auto-scroll. Instead:
> 1. I add the new message to the bottom of the messages array (state update)
> 2. I increment the 'new messages' count on the floating scroll-to-bottom button
> 3. The button shows something like '↓ 3 new messages'
> 4. When the user clicks it, I smooth-scroll to the bottom and reset the count
>
> This is exactly what WhatsApp and Slack do."

> **Interviewer: "How would you handle jumping to a specific message? Like when someone clicks on a message in search results."**

> "This is called **scroll-to-index** or **anchor scrolling**. Steps:
> 1. If the message is already loaded in memory, find its index and use the virtualizer's `scrollToIndex()` method
> 2. If it's not loaded (very old message), I need to fetch messages around that timestamp from the API, load them into state, then scroll to the target
> 3. After scrolling, briefly **highlight** the target message (yellow flash animation) so the user knows which one it is
>
> This is one of the harder features to implement well with virtualization, because you may need to jump to an arbitrary position in a list of unknown total length."

---

## Step 9: Polish & Tradeoffs (Minutes 48-55)

### Mention These Unprompted (They Impress Interviewers)

#### 1. Accessibility
> "For the chat app to be accessible:
> - The message input has proper `aria-label`
> - New messages are announced to screen readers using an `aria-live='polite'` region
> - The conversation list and messages are navigable with keyboard (arrow keys + Enter)
> - The 'X new messages' button is focusable and announced
> - Color contrast for read/unread indicators meets WCAG AA"

#### 2. Security
> "Messages should be sanitized before rendering to prevent XSS. If I'm rendering HTML content (e.g., links), I use a sanitizer like DOMPurify. Ideally, messages are plain text rendered in a text node, not `innerHTML`.
>
> For sensitive chats, we could implement **end-to-end encryption** on the client using the Web Crypto API (Signal Protocol). But that's a significant scope increase."

#### 3. Performance Budget
> - "Initial load: < 200KB JS (code-split the emoji picker, settings, etc.)"
> - "Message render: < 16ms per frame (60fps scrolling)"
> - "Time to first message visible: < 2s (use cached data from IndexedDB)"

#### 4. Monitoring
> "In production, I'd track:
> - WebSocket connection drops and reconnection frequency
> - Message delivery latency (time from send to ACK)
> - Failed message rate
> - Client-side JS errors (Sentry)
> - Core Web Vitals (LCP, INP, CLS)"

### Tradeoffs Summary

| Decision | Choice | Alternative | Why I Chose This |
|----------|--------|-------------|------------------|
| Real-time transport | WebSocket | SSE + REST | Bidirectional, low overhead for chat |
| State management | Zustand | Redux, Context | Lightweight, selector-based re-renders |
| Offline storage | IndexedDB | localStorage | Async, larger storage, structured data |
| Message ordering | Server timestamp | Client timestamp | Consistent ordering across all clients |
| Virtualization | @tanstack/virtual | react-window | Better dynamic height support |
| ID generation | Client UUID | Server-assigned | Enables optimistic UI, idempotent retries |

---

## Cross Questions Bank (with Answers)

Here are **25 cross-questions** interviewers commonly ask, grouped by topic. Each has a concise answer.

### Architecture

**Q1: "How would you extend this to group chat?"**
> "Two key changes:
> 1. Data model: `Conversation` gets `type: 'group' | 'direct'` and the `participants` array can have N members
> 2. Read receipts become more complex -- I need to track read status per participant. Instead of double-blue-check, I might show '3 of 8 read' or show individual avatars like WhatsApp groups
> 3. Typing indicators show who is typing: 'Alice is typing...' or 'Alice and Bob are typing...'
>
> The WebSocket message format stays the same -- the server handles fan-out to all group members."

**Q2: "How would you add message reactions (emoji reactions like Slack)?"**
> "Add a `reactions` field to the Message model: `Record<string, string[]>` (emoji → [userIds]). Clicking a reaction sends a WebSocket event `{type: 'reaction', messageId, emoji}`. The server broadcasts to all participants. On the UI, I show reaction pills below the message bubble with counts."

**Q3: "What if the server goes down completely? What does the user experience?"**
> "The WebSocket closes, triggering reconnection attempts with exponential backoff. The UI shows 'Connecting...' banner. The user can still:
> - Read all cached conversations and messages (from IndexedDB)
> - Type and send messages (queued locally)
> - Browse the app normally
>
> When the server comes back, everything syncs. The user might see brief reordering as server timestamps are assigned."

**Q4: "Would you use REST or GraphQL for the non-real-time API?"**
> "For a chat app, I'd lean REST:
> - The data model is relatively simple (conversations, messages, users)
> - REST's caching with HTTP cache headers works well for message history
> - GraphQL's flexibility isn't needed here since the queries are predictable
>
> If this were a more complex app with many interrelated entities (like a social network), GraphQL's ability to fetch exactly what you need in one request would be more valuable."

### WebSocket

**Q5: "What if WebSocket is blocked by a corporate firewall?"**
> "Fallback to **long polling**: the client makes a GET request that the server holds open until there's new data (or a timeout of ~30s). When data arrives or timeout hits, the server responds, and the client immediately makes a new request. Libraries like Socket.IO handle this fallback automatically."

**Q6: "How do you scale WebSocket servers horizontally?"**
> "This is more of a backend concern, but from the frontend perspective, I need to know that my WebSocket connection might be to a different server after reconnect. So:
> - Messages must be identified by unique IDs, not connection state
> - On reconnect, I send a `sync` request with my last-received timestamp
> - The server (using Redis pub/sub or similar) ensures any server can serve any user"

**Q7: "What protocol would you use on top of WebSocket?"**
> "Plain JSON for simplicity. Each message is `{type: string, payload: any}`. For a production app at scale, I'd consider Protocol Buffers (smaller payload, strict schema) or MessagePack (binary JSON, ~30% smaller). But JSON is debuggable in DevTools, which matters during development."

### Messages

**Q8: "How do you prevent duplicate messages?"**
> "Client-generated UUIDs are the key. Every message gets a UUID before sending. The server uses this as an idempotency key -- if it receives the same ID twice (retry scenario), it ignores the duplicate and just re-sends the ACK. On the client, I also deduplicate by ID when inserting into the messages array."

**Q9: "How would you implement message search?"**
> "This depends on scale:
> - **Simple approach**: API endpoint `GET /messages?q=search_term&conversationId=xyz` with server-side full-text search (Elasticsearch or PostgreSQL full-text)
> - **Client-side**: For the active conversation, I could search through cached messages in IndexedDB for instant results, and show server results for full history
> - **UI**: A search bar that shows results as a list. Clicking a result jumps to that message in the conversation (scroll-to-message feature)"

**Q10: "How would you handle message editing and deletion?"**
> "For editing: the server sends a `{type: 'message_edited', messageId, newContent, editedAt}` event via WebSocket. I update the message in state and show '(edited)' text. I'd keep the original content in the database for audit purposes.
>
> For deletion: `{type: 'message_deleted', messageId}`. I replace the message content with 'This message was deleted' (like WhatsApp) rather than removing it from the array, to avoid confusing gaps in conversation."

### Performance

**Q11: "The conversation list shows 500 conversations. How do you keep it fast?"**
> "Virtualize the conversation list too. Only render the ~15-20 visible conversations. The rest exist only in state/memory. Scrolling renders them on demand. Each `ConversationItem` is wrapped in `React.memo` since they rarely change."

**Q12: "What if a user is in 100 group chats and all are active?"**
> "Prioritize updates:
> 1. Active conversation: full real-time updates (messages, typing, presence)
> 2. Other conversations: only update `lastMessage` and `unreadCount` (lightweight)
> 3. Batch non-active conversation updates: instead of re-rendering the list on every message, batch updates with a 500ms debounce
>
> Also, I wouldn't maintain WebSocket subscriptions for all groups simultaneously. I'd subscribe to the active conversation for detailed events and receive only notification-level events for others."

**Q13: "How do you handle images and media in messages?"**
> "Multi-step process:
> 1. User selects an image → show a local preview immediately (URL.createObjectURL)
> 2. Upload to cloud storage (S3/CloudFront) via REST API with progress tracking
> 3. Once uploaded, send the message via WebSocket with the media URL
> 4. Receiver gets the URL, renders a thumbnail (server-generated), click to view full
>
> For performance: lazy load images not in viewport, use srcset for different resolutions, show a blurred placeholder (blur hash) during load."

### Offline & Sync

**Q14: "How much data do you cache offline?"**
> "Bounded caching:
> - Last 20 conversations (metadata)
> - Last 50 messages per cached conversation
> - All pending outgoing messages (unbounded, but realistically small)
>
> Total: maybe 2-5 MB. Well within IndexedDB limits. I'd also implement LRU eviction -- if the cache grows beyond a threshold, drop the oldest conversations."

**Q15: "What happens if the user was offline for 3 days?"**
> "On reconnect, the sync could return thousands of messages. I wouldn't load all at once:
> 1. First, sync conversation metadata (who sent what, unread counts) -- this is lightweight
> 2. For the conversation the user opens, fetch recent messages on demand
> 3. Background-sync other conversations gradually
>
> Show a loading indicator during initial sync: 'Syncing messages...'"

### Security

**Q16: "How do you store auth tokens on the client?"**
> "Short answer: **httpOnly cookies** for the session, or an **in-memory variable** for the access token.
>
> Never in localStorage (vulnerable to XSS). An httpOnly cookie can't be read by JavaScript, so it's safe from XSS. For the WebSocket auth, I'd use a short-lived token fetched via REST (which sends the httpOnly cookie automatically) and pass it during WebSocket handshake."

**Q17: "How do you prevent XSS in chat messages?"**
> "Never use `dangerouslySetInnerHTML` for message content. Render messages as text nodes (`{message.content}` in JSX automatically escapes HTML).
>
> If I need to render rich content (links, mentions), I'd parse the text and create React elements: detect URLs with a regex and wrap them in `<a>` tags, detect @mentions and wrap them in `<span>` with styling. But the raw text is never injected as HTML."

### UX & Edge Cases

**Q18: "What if the user has slow internet (2G)?"**
> "Adaptive UX:
> - Detect connection quality using `navigator.connection.effectiveType`
> - On slow connections: disable auto-loading images, reduce WebSocket heartbeat frequency, batch multiple small messages into single payloads
> - Show clear loading states and progress indicators
> - Offline queue becomes even more important -- messages might sit in 'sending' for a while"

**Q19: "How does the typing indicator not become annoying?"**
> "Two safeguards:
> 1. **Throttle on sender side**: only send typing event every 3 seconds, not every keystroke
> 2. **Timeout on receiver side**: hide typing indicator after 4 seconds of no new typing event. This handles the case where the user stopped typing but the 'stopped_typing' event was lost
>
> Also, don't show typing indicator for the user's own messages, and in group chats, show at most 2 names ('Alice and Bob are typing...' or 'Alice and 3 others are typing...')"

**Q20: "How would you implement push notifications?"**
> "Using the **Push API** and **Service Worker**:
> 1. On app load, request notification permission
> 2. Register a push subscription with the server (subscription endpoint from the browser)
> 3. When a message arrives and the user is NOT on the page, the server sends a push notification via the browser's push service
> 4. The Service Worker receives it and shows a native notification using the Notification API
>
> When the user clicks the notification, the Service Worker opens/focuses the app and navigates to the correct conversation."

**Q21: "How do you handle the user having the same conversation open on phone and laptop?"**
> "The server handles this by sending messages to all active connections for that user. Both devices get the message, both update their UI. Read receipts from either device sync to the other -- if the user reads on phone, the laptop should also clear the unread count.
>
> On reconnect, each device syncs independently using the last-received timestamp. The result is eventually consistent."

**Q22: "What about link previews (like Slack shows a preview card when you paste a URL)?"**
> "This is **unfurling**:
> 1. When a message is sent, the server detects URLs in the content
> 2. The server fetches Open Graph metadata (title, description, image) from the URL
> 3. This metadata is attached to the message and sent to clients
> 4. The client renders a preview card below the message text
>
> Important: the **server** fetches the URL, not the client. Fetching from the client would leak the user's IP and could be a security risk (SSRF in reverse). The server can also cache unfurled data."

**Q23: "How would you test this chat application?"**
> "Testing pyramid:
> - **Unit tests**: message ordering logic, timestamp formatting, status calculations (Jest)
> - **Component tests**: MessageBubble renders correctly for each status, ConversationList sorts properly (React Testing Library)
> - **Integration tests**: full send/receive flow with a mocked WebSocket (MSW or custom WebSocket mock)
> - **E2E tests**: two browser instances sending messages to each other (Playwright)
>
> I'd particularly focus on testing edge cases: offline → online transitions, out-of-order messages, duplicate messages, reconnection scenarios."

**Q24: "What if we need end-to-end encryption?"**
> "E2E encryption means the server can't read message content. On the frontend:
> - Use the Web Crypto API to generate key pairs per user
> - Exchange public keys during conversation setup (via the server, but it can't decrypt)
> - Encrypt each message client-side before sending, decrypt on receipt
> - This breaks server-side search and link unfurling (server can't read content)
>
> This is conceptually the Signal Protocol. It's a massive scope increase, so I'd mention it as a future consideration."

**Q25: "You mentioned you'd use Zustand. Walk me through how you'd structure the store."**
> ```typescript
> const useChatStore = create((set, get) => ({
>   conversations: {},
>   messages: {},
>   activeConversationId: null,
>
>   addMessage: (conversationId, message) => set(state => ({
>     messages: {
>       ...state.messages,
>       [conversationId]: insertMessageInOrder(
>         state.messages[conversationId] || [],
>         message
>       ),
>     },
>   })),
>
>   updateMessageStatus: (messageId, status) => set(state => {
>     // find and update the message across conversations
>   }),
>
>   setActiveConversation: (id) => set({ activeConversationId: id }),
> }));
> ```
>
> "Components subscribe to specific slices using selectors, so only relevant components re-render:
> `const messages = useChatStore(state => state.messages[conversationId]);`"

---

## Cheat Sheet: Things That Impress Interviewers

These are things most candidates DON'T mention. Bring them up naturally.

### 1. Say "trade-off" a lot (genuinely)
Don't just pick a technology -- explain what you'd lose.
> "I'd go with WebSocket over SSE. The trade-off is that WebSocket requires more server infrastructure and doesn't work through some proxies, but for chat, bidirectional communication outweighs those costs."

### 2. Mention failure scenarios unprompted
> "What happens if the WebSocket connection drops mid-message-send? The message is in our queue. On reconnect, we re-send with the same ID so the server can deduplicate."

### 3. Show awareness of real products
> "WhatsApp uses a similar approach for read receipts -- they send a watermark (last read message ID) rather than individual receipts per message. This reduces traffic significantly."

### 4. Mention metrics you'd track
> "In production, I'd monitor message delivery latency as a p50/p95 metric. If p95 goes above 1 second, something's wrong with our WebSocket infrastructure."

### 5. Acknowledge what you don't know
> "I'm not deeply familiar with CRDT algorithms for conflict resolution, but I know that's the approach Figma uses for real-time collaboration. For our chat app, server-ordered timestamps are sufficient since messages don't conflict."

This is 10x more impressive than bluffing.

### 6. Tie back to user experience
> "The reason I save scroll position is purely UX -- nothing is more frustrating than losing your place in a conversation because new messages pushed you around."

### 7. Think about the empty states and loading states
> "When the user first opens a conversation, I show skeleton message bubbles while loading. If the conversation has no messages, I show a friendly empty state: 'Say hi to Alice!' with a wave emoji."

### 8. As someone with basic frontend experience specifically:

**What to lean into:**
- Be honest about your experience level. Saying "In my experience with simpler apps, I've handled X, and I'd approach this at scale by doing Y" is genuine and credible
- Focus on **fundamentals** you know well (React state management, component design, CSS layout)
- Show curiosity: "I haven't implemented WebSocket reconnection with exponential backoff in production, but here's how I'd approach it based on what I know about retry patterns"

**What to avoid:**
- Don't pretend you've built WhatsApp before
- Don't use buzzwords you can't explain if probed
- Don't rush past areas you're weak in -- slow down and reason through them

**A strong signal for a mid-level candidate is:**
> "I know what I know, I know what I don't know, and I can reason through the things I don't know."

---

## Summary: The 5-Minute Recap

If you only remember 5 things about this design:

1. **WebSocket for real-time, REST for history** -- two communication channels, each for the right job
2. **Optimistic UI with client-generated IDs** -- show the message immediately, reconcile with server timestamp later, use the same ID for idempotent retries
3. **Exponential backoff with jitter for reconnection** -- prevents thundering herd, handles flaky networks gracefully
4. **Scroll position management** -- auto-scroll if near bottom, preserve position if scrolled up, restore position when loading older messages above
5. **IndexedDB for offline queue** -- survives tab closes, enables offline-first experience, bounded cache with LRU eviction

Practice explaining each of these in 2 minutes. If you can do that fluently, you'll do well.

---

*Estimated study time: 3-4 hours to read and understand, then 2-3 practice runs explaining it aloud (use a timer, 45 minutes each).*
