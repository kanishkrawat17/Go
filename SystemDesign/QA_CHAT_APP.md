# Chat App System Design -- Q&A Learnings from Practice Session

> These are all the questions I asked, confusions I had, and clarifications I got
> while studying the chat app system design. This is the "what I actually learned"
> companion to SYSTEM_DESIGN_CHAT_APP.md.

---

## 1. How to study this effectively

- **Layer 1** (Day 1): Read for understanding, focus on WHY each decision is made (~2 hrs)
- **Layer 2** (Day 2): Close the doc, explain each section OUT LOUD to a wall (~1.5 hrs)
- **Layer 3** (Day 3): Draw the entire system from memory on paper (~1 hr)
- **Layer 4** (Day 4): Practice cross-questions -- read question, cover answer, try yourself (~1.5 hrs)
- **Layer 5** (Day 5, 10, 14): Full 45-min mock interview simulation (~1 hr each)
- Total: ~10 hours spread over 2 weeks. Spaced repetition is key.

---

## 2. What other system designs does the chat app cover?

The chat app teaches 6 patterns that cover ~80% of all frontend system design questions:

| Pattern | Reappears in |
|---|---|
| WebSocket + real-time | Notifications, live scoreboard, stock ticker, collaborative editing |
| Optimistic UI + idempotent retries | Twitter likes, Kanban drag-drop, comments, cart |
| Infinite scroll + scroll management | Any feed (Twitter, Instagram, news), email, comments |
| Offline + IndexedDB + sync | Notes app, todo app, Google Docs, maps |
| Virtualization | Any large list, data tables, spreadsheets, file explorers |
| Normalized state (Record<id, T>) | Literally every frontend system design |

**Best next topics after chat app:** Twitter Feed + Google Docs. Together with chat, covers ~85% of questions.

---

## 3. Zustand vs IndexedDB -- When does each matter?

This was my biggest confusion. Here's the clarity:

```
Zustand  = RAM (browser memory)
           → FAST, but dies when you close the tab or refresh
           → Has NOTHING to do with internet. Going offline does NOT kill Zustand.

IndexedDB = Disk (browser storage)
           → Slower, but survives tab close, refresh, browser restart
           → Also has nothing to do with internet. It's a local file.

Server   = Remote database
           → Needs internet to reach
```

**Scenario A: Go offline, keep tab open, come back online**
- Zustand handles everything. IndexedDB not needed.
- UI stays exactly the same. Just a "Reconnecting..." banner appears.

**Scenario B: Go offline, CLOSE the tab, reopen later**
- Zustand dies (RAM cleared). IndexedDB still has the data (disk).
- On app startup: persist middleware reads from IndexedDB → fills Zustand → UI works.
- No internet needed to read from IndexedDB.

**How the auto-sync works:**
- Zustand's `persist` middleware automatically writes to IndexedDB on every state change.
- On app startup, it automatically reads from IndexedDB to rehydrate Zustand.
- You never manually save or load. It just works.

**Key interview line:**
> "Zustand is the runtime source of truth. The persist middleware mirrors state to IndexedDB
> on every update. IndexedDB is only read on app startup to rehydrate after a tab close."

---

## 4. Does it have to be perfect in an interview?

NO. What matters:

| What they want | What they DON'T care about |
|---|---|
| You know WHICH events exist and WHY | Exact JSON payload field names |
| You understand the DATA FLOW | Whether you said "status" vs "presence" |
| You can REASON through gaps | Having every detail memorized |

**The 6 concepts to say fluently (even in your own words):**
1. "Message is shown immediately (optimistic), server confirms later"
2. "Messages go: sending → sent → delivered → read"
3. "Read receipts use watermark pattern (one event, not one per message)"
4. "Typing indicators are throttled, ephemeral, auto-expire"
5. "Online status is a Set for fast lookup"
6. "Offline messages queue and flush on reconnect, same ID prevents duplicates"

---

## 5. My data model approach -- is it valid?

I tried my own structure instead of the textbook one:

```
conversations = {
  "1": { userId, id, participants, username, status,
         profileThumbnailUrl, unreadMsgsCount,
         lastMessageTimeStamp, lastUnreadMessageContent },
}
```

**Verdict: VALID.** Interviewer would pass this. The one probe they'd make:
- "username and profileThumbnailUrl are denormalized here. If user changes their pic,
   you'd update it in every conversation. Consider extracting to a separate users lookup."

**Strong response:** "For simplicity I'm denormalizing. If needed, I'd extract to
a separate `users: Record<id, User>` and just store participantIds here."

---

## 6. My MessageListing approach -- is it valid?

```
MessageListing = {
  conversationId,
  typingIndicator,
  status,                // other person's online status
  messages: Message[],   // last 10, load more on scroll up with cursor pagination
}
```

**Verdict: VALID.** Co-locating conversation data is a legitimate pattern.

**What to add:** `hasMore: boolean` and `oldestLoadedTimestamp` for cursor pagination.
**What to add to Message:** `senderId` (need it for left/right bubble alignment).
**Strong line:** "I'm only loading the last 10 messages upfront. When the user scrolls up,
I hit the API with cursor-based pagination using the oldest loaded timestamp."

---

## 7. Cursor-based pagination -- how it works concretely

```
8:00 AM  Open chat    → GET /messages?limit=10              → get msg 91-100
8:02 AM  Scroll up    → GET /messages?before=7:30AM&limit=10 → get msg 81-90
8:03 AM  Scroll again → GET /messages?before=11PM&limit=10   → get msg 71-80
...keep going until server returns < 10 messages → hasMore = false
```

Why cursor > page numbers: "In a chat app, new messages arrive constantly.
Page 3 today shows different messages than page 3 did 5 minutes ago.
Cursor-based (before=timestamp) is stable regardless of new messages."

---

## 8. "New messages while scrolled up" -- where does that state live?

My instinct: add `newMessageWhileOpenAndNotInView` to the data model.

**Better approach:** It's just a `useState(0)` counter (unseenCount) INSIDE the
MessageList component. Not in the global store.

**Why local, not global:**
- Only MessageList needs this count
- It resets when you switch conversations (component unmounts)
- Putting it in global store adds noise for no benefit

**Rule of thumb:**
> "Global store = shared across components or survives navigation.
> Local useState = only one component needs it."

---

## 9. Message retry -- how it actually works

DON'T add a separate `isFailed: boolean`. The existing `status: 'failed'` covers it.
Adding both creates two sources of truth (what if status='sent' but isFailed=true?).

**The flow:**
1. Send message → status = 'sending' ⏳ → start 10s timeout
2. No ACK in 10s → status = 'failed' ❌ → show retry button
3. User taps retry → status = 'sending' ⏳ → re-send with SAME message ID
4. Same ID is key: server deduplicates if original DID arrive but ACK was lost

**Interview line:**
> "Failed messages show a retry button. On tap, I re-send with the same message ID
> so the server can deduplicate. The status goes back to 'sending' and I restart
> the ACK timeout."

---

## 10. What "data modeling" means in frontend system design

```
Backend data modeling  = designing DATABASE TABLES (tables, columns, foreign keys)
Frontend data modeling = designing the CLIENT STORE (state shape, components, connections)
```

The flow:
1. Look at the UI → "sidebar shows conversations with name, preview, badge"
2. What data does this need? → `conversations: Record<id, Conv>` with lastMessage, unreadCount
3. What component renders it? → `<ConversationItem>` takes conversation as prop
4. Global or local? → Global (sidebar + chat window both need conversations)

That's it. You're mapping UI → data → components → global vs local.

---

---

## 11. Auth for WebSocket (related it to my own workflow-apis repo)

My own app uses **cookie-based auth**: browser sends `auth` cookie automatically with
`credentials: 'include'` in fetch calls. Server validates via user-token-validator service.

WebSocket CAN'T do this (no header/cookie support in browser WebSocket API). So:
1. Get a short-lived token via REST (which has cookie auth)
2. Send that token as the first WebSocket message: `{type: 'auth', token: '...'}`
3. Server rejects all other messages until auth succeeds

**Interview line:**
> "WebSocket doesn't support headers. I'd get a short-lived token via REST,
> then send it as the first WebSocket message. Server rejects everything until auth succeeds."

---

## 12. User B is offline when A sends a message

Server acts as a **store-and-forward mailbox**:
- Server receives A's message → saves to DB → ACKs to A (✓ sent)
- Tries to forward to B → B has no active WebSocket → holds in DB
- When B comes online → B sends sync request → server delivers all missed messages
- B sends delivery_ack → server forwards to A → A sees ✓✓

Message stays at ✓ (sent) until B comes online. Could be minutes or days.

---

## 13. Persistent DB -- when does it get updated?

**Server saves to DB FIRST, before ACKing or forwarding. Always.**

```
Message arrives at server → Save to DB → THEN ACK to sender + forward to receiver
```

Why? If server crashes after forwarding but before saving, message is lost forever.
DB first = message is safe even if everything else fails.

**What gets saved vs what doesn't:**
- Messages, read status, delivery status → YES (user expects to see these tomorrow)
- Typing indicators, presence, heartbeat → NO (ephemeral, in-memory only)

**Rule:** "Will the user expect to see this data tomorrow? YES = DB. NO = memory only."

---

## 14. Optimistic UI -- showing ⏳ not ✓

Optimistic UI = show the MESSAGE immediately. Not the STATUS.
- Show the message in the chat right away (optimistic)
- Show ⏳ not ✓ (honest -- server hasn't confirmed yet)
- Showing ✓ immediately would be lying to the user

**Optimistic = "I'll show the message."  Honest = "I'll show the real status."**

---

## 15. Offline: ⏳ vs ❌ -- when to show which

```
Online + no ACK in 10s  → ❌ error (something actually broke, let user retry)
Offline                 → ⏳ queued (nothing broke, just waiting for internet)
```

Don't show ❌ when offline -- user knows they're offline, errors are just noise.
Show ⏳ + banner "Messages will send when you're back online."
Messages auto-send on reconnect. User doesn't tap anything.

---

## 16. Typing indicator -- throttle not debounce

- SENDER: throttle every 3 seconds (send at most 1 event per 3s)
- RECEIVER: show "typing..." + 4-second timeout. Reset timer on each new event.
- No ACK needed. Fire and forget.
- Why throttle not debounce? Debounce waits until user STOPS typing. That would
  delay showing the indicator for the entire typing duration. Throttle shows it immediately.

---

## 17. Reconnect order: receive first, send second

On reconnect after being offline:
1. First sync missed messages FROM server (receive)
2. Then flush pending queue TO server (send)

Why? So our pending messages get timestamps AFTER the messages we missed.
Keeps ordering correct.

---

## 18. My architecture diagram feedback

Drew the architecture on a whiteboard. Got all boxes right:
UI Layer → Data Layer (Zustand) ← WS Manager → IndexedDB, WS Server, REST API

Fixes needed:
- UI ↔ Data Layer should be bidirectional (reads AND dispatches)
- WS Manager needs arrow to/from WS Server (currently floating)
- REST API responses should go through Data Layer (not directly to UI)

---

---

## 19. Message lifecycle diagram -- my attempt and feedback

Drew a sequence diagram on whiteboard (User A → WS Server → User B). Got 3 out of 4 steps right.

**What I missed:** The server ACK back to User A (step 2). Without this, A doesn't know
if the server even got the message. The status would stay at ⏳ forever.

**Correct flow (4 steps, not 3):**
```
A sends message → Server saves to DB → Server ACKs to A (✓) + forwards to B
→ B sends delivery_ack → Server forwards to A (✓✓)
```

**Key: always include message ID in every payload** so you know WHICH message each event refers to.

**Score if presented in interview:** 7/10 (passing). Interviewer would nudge about the missing ACK,
I'd correct immediately. That's normal whiteboard interaction.

---

## 20. The WS Server sits between User A and User B

```
❌ NOT how it works:    User A ────────► User B     (peer-to-peer)
✅ How it works:        User A ──► Server ──► User B (server in the middle)
```

Users NEVER talk directly. Every message goes through the server.
Both users have their own persistent WebSocket connection to the same server.
Server knows which connection belongs to which user and routes messages accordingly.

---

## 21. Server always saves to DB FIRST

```
Message arrives → Save to DB → THEN ACK + forward
```

Why DB first? If server ACKs and forwards but crashes before saving:
- A thinks "sent ✓"
- B sees the message
- Database has nothing
- On refresh, message is GONE. Data loss.

DB first = message is safe even if server crashes in the next millisecond.

---

## 22. Optimistic UI: show MESSAGE immediately, not STATUS

```
Optimistic ≠ lying.

Show the message instantly  → YES (optimistic -- don't wait to display it)
Show ✓ (sent) instantly     → NO (that's lying -- server hasn't confirmed)
Show ⏳ (sending) instantly → YES (honest -- we're trying, not confirmed yet)
```

Restaurant analogy:
- Optimistic: "Got it, coming right up!" (order placed, kitchen working)
- Lying: "Your food is ready!" (kitchen hasn't even seen the order)

---

## 23. Offline: show ⏳ (queued) not ❌ (error)

```
Online + no ACK in 10s  → ❌ (real failure, show retry button)
Offline                 → ⏳ (just waiting, will auto-send on reconnect)
```

Don't show 5 error icons when user is on the subway. They know they're offline.
Show ⏳ + banner "Messages will send when you're back online."
All queued messages auto-send on reconnect. No user action needed.

Check connection status FIRST before deciding ❌ vs ⏳.

---

## 24. Typing indicator -- throttle (not debounce), fire-and-forget (no ACK)

- SENDER: throttle 1 event per 3 seconds (not every keystroke)
- RECEIVER: show "typing..." + auto-clear after 4 seconds of no new event
- NO ACK needed. If a typing event is lost, nobody cares.
- Why throttle not debounce? Debounce waits until user STOPS typing → other person
  sees nothing until typing is done. Defeats the purpose. Throttle shows it immediately.

---

## 25. Reconnect after being offline: receive FIRST, send SECOND

```
1. Reconnect
2. Sync missed messages FROM server (receive)    ← first
3. THEN flush our pending queue TO server (send)  ← second
```

Why this order? Our pending messages get server timestamps AFTER the messages we missed.
Keeps the chronological ordering correct.

---

## 26. Unseen message count while scrolled up -- local state, not global

When user is scrolled up and new messages arrive, show "↓ 3 new messages" button.

This is `useState(0)` inside MessageList component. NOT in Zustand.
- Only MessageList needs it
- Resets on conversation switch (component unmounts)
- Increment when `messages.length` increases AND `isNearBottom` is false
- Reset to 0 when user scrolls to bottom or clicks the button
- Use `useRef` to track previous message count to know how many arrived

---

## 27. Scroll position restoration on loading older messages

**The problem:** When you prepend 20 older messages above, the browser keeps `scrollTop`
at the same pixel value. But the content shifted down, so user suddenly sees message 1
instead of message 21 where they were.

**The fix (3 lines):**
```
const oldScrollHeight = container.scrollHeight;     // save before
prependMessages(olderMessages);                      // add to state
requestAnimationFrame(() => {
  const newScrollHeight = container.scrollHeight;    // measure after
  container.scrollTop = newScrollHeight - oldScrollHeight;  // adjust
});
```

**Book analogy:** You're reading page 5. Someone adds 20 pages at the beginning.
Your page 5 is now page 25. Move your finger from position 5 to position 25. Same content.

---

## 28. Reply-to-message feature -- data model addition

```
Message {
  id, content, status, timestamp, senderId,
  repliedMsgId,     // ← null if not a reply, message ID if it is
}
```

When user taps on a reply preview:
- **If message is loaded:** scroll to it + highlight with yellow flash
- **If message is not loaded (old):** call API `GET /messages?around=repliedMsgId&limit=30`,
  load into state, scroll to target, highlight
- **If message was deleted:** show "Original message was deleted", tap does nothing

---

## 29. Three layers of storage -- when each is used

```
PERSISTENT DB (server)          INDEXEDDB (client disk)         ZUSTAND (client RAM)
──────────────────────         ──────────────────────          ──────────────────────
The TRUTH.                     A CACHE.                        The LIVE state.
Every message ever.            Last 50 messages.               What's on screen now.

Written by: server             Written by: persist             Written by: WS Manager
                               middleware (auto)               + REST responses
Read by: REST API              Read on: app startup            Read by: React components

Survives: everything           Survives: tab close,            Survives: nothing
                               browser restart                 (dies on refresh/close)

Used when: refresh, history,   Used when: app startup          Used when: app is running
new device                     (instant load before APIs)      (every second)
```

---

## 30. My current readiness (final update)

- Concepts: ✅ Strong. Understand all core patterns deeply.
- Data modeling: ✅ Can design my own shapes and explain tradeoffs.
- Going deep on probes: ✅ Can reason through follow-ups in real-time.
- Communication: 🔶 Getting better. Need mock practice to polish delivery.
- Architecture diagram: ✅ Can draw from memory (with minor fixes).
- Message lifecycle: ✅ Can draw the sequence diagram (missed 1 step, self-corrected).
- Scroll management: ✅ Understand auto-scroll, infinite scroll, position restoration.
- **Progress:** Covered Steps 1-9 of the system design doc + 30 Q&A learnings.
- **Remaining:** Step 10 (tradeoffs wrap-up) + Cross Questions practice.
- **Action:** Finish Step 10 (10 min), practice cross questions (30 min), then do 3 mock runs.

---

## 31. Files I have for this prep

| File | What it is |
|---|---|
| `FRONTEND_INTERVIEW_PREP.md` | Master 16-week prep plan (all topics) |
| `SYSTEM_DESIGN_CHAT_APP.md` | Full chat app system design (interview walkthrough) |
| `CHAT_APP_QA_LEARNINGS.md` | This file -- 30 Q&A learnings from practice sessions |

---

*Last updated: March 15, 2026*
