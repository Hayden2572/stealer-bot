<p align="center">
  <img src="https://img.shields.io/badge/DIRHUNTER-STEALER%20MODE-black?style=for-the-badge&logo=linux&logoColor=red" />
</p>

<h1 align="center">⚡ DIRHUNTER — Aggressive Async Directory Scanner & Telegram Uploader ⚡</h1>

DirHunter is a **fast, concurrent, root-level directory hunter** that locks onto a target folder, rips it into a ZIP archive, and launches it straight into your Telegram.  

---

## 🔥 Features (No BS)

- ⚡ **Aggressive recursive scan** starting at `/`
- 🎯 Dual-stage match system:  
  - `flag` → folder name to hit  
  - `subFlag` → optional path filter  
- 💥 **Auto-terminate scan** the moment a match is detected
- 📦 ZIP archive packing via `archiver/v4`
- 🚀 **Direct fire** into Telegram via Bot API
- 🧵 Massive async traversal (goroutines + WaitGroup)
- 🧱 Scan-safe architecture (RWMutex + visited map)

---

## ⚠️ Requirements & Limitations (Read Before Using)

| Requirement | Meaning |
|------------|---------|
| 🌐 Internet | Absolutely required for Telegram upload |
| 📁 50 MB cap | Telegram refuses anything larger |
| 🔐 Permissions | Must be able to read scanned directories |

---

## 🛠 Configuration

All key variables live in the **data section** of the code:

| Variable | What It Does |
|----------|--------------|
| `flag` | Target folder name |
| `subFlag` | Forces match only inside paths containing this substring |
| `startPath` | Root of the scan (`/` by default) |
| `botToken` | Telegram bot token |
| `chatId` | Target chat for upload |
| `done` | Scan-stop switch |
| `visited` | Map preventing re-entry to same paths |
| `resultDir` | Automatically set once the match is found |

---

## 🧬 How It Operates

1. Launches a full-depth async crawl of the filesystem.
2. When a directory name equals `flag`:
   - If `subFlag` is set → validates that path contains it  
   - If valid → locks onto target
3. Instantly aborts all further scanning.
4. Compresses the captured directory into `data.zip`.
5. Fires the archive to Telegram’s `sendDocument` endpoint.

---

## ▶️ Run

```bash
go mod tidy
go run .

## ⚠️ Legal Notice

Use with extreme caution and responsibility

    🚫 Unauthorized access to computer systems is illegal

    🔒 Obtain proper permissions before scanning any system

    👮 You are solely responsible for your actions

    ✅ Intended for authorized security testing only

If you don't own it - don't scan it
