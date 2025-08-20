import { useState, useRef, useEffect } from 'react'
import './App.css'

interface ChatMessage {
  id: string
  text: string
  sender: 'user' | 'jarvis'
  timestamp: Date
}

function App() {
  const [messages, setMessages] = useState<ChatMessage[]>([])
  const [inputText, setInputText] = useState('')
  const [isLoading, setIsLoading] = useState(false)
  const messagesEndRef = useRef<HTMLDivElement>(null)

  const scrollToBottom = () => {
    messagesEndRef.current?.scrollIntoView({ behavior: 'smooth' })
  }

  useEffect(() => {
    scrollToBottom()
  }, [messages])

  const sendMessage = async () => {
    if (!inputText.trim() || isLoading) return

    const userMessage: ChatMessage = {
      id: Date.now().toString(),
      text: inputText,
      sender: 'user',
      timestamp: new Date()
    }

    setMessages(prev => [...prev, userMessage])
    setInputText('')
    setIsLoading(true)

    try {
      const response = await fetch('http://localhost:7070/chat', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({ message: inputText }),
      })

      if (!response.ok) {
        throw new Error(`HTTP error! status: ${response.status}`)
      }

      const data = await response.json()

      if (data.error) {
        throw new Error(data.error)
      }

      const jarvisMessage: ChatMessage = {
        id: (Date.now() + 1).toString(),
        text: data.response,
        sender: 'jarvis',
        timestamp: new Date()
      }

      setMessages(prev => [...prev, jarvisMessage])
    } catch (error) {
      console.error('Error sending message:', error)
      const errorMessage: ChatMessage = {
        id: (Date.now() + 1).toString(),
        text: `Error: ${error instanceof Error ? error.message : 'Failed to send message'}`,
        sender: 'jarvis',
        timestamp: new Date()
      }
      setMessages(prev => [...prev, errorMessage])
    } finally {
      setIsLoading(false)
    }
  }

  const handleKeyPress = (e: React.KeyboardEvent) => {
    if (e.key === 'Enter' && !e.shiftKey) {
      e.preventDefault()
      sendMessage()
    }
  }

  return (
    <div className="chat-container">
      <header className="chat-header">
        <h1>Jarvis Chat</h1>
        <p>Chat with your AI assistant</p>
      </header>

      <div className="chat-messages">
        {messages.length === 0 ? (
          <div className="welcome-message">
            <p>Welcome! Start chatting with Jarvis by typing a message below.</p>
          </div>
        ) : (
          messages.map((message) => (
            <div key={message.id} className={`message ${message.sender}`}>
              <div className="message-content">
                <div className="message-text">{message.text}</div>
                <div className="message-time">
                  {message.timestamp.toLocaleTimeString()}
                </div>
              </div>
            </div>
          ))
        )}
        {isLoading && (
          <div className="message jarvis">
            <div className="message-content">
              <div className="message-text typing">Jarvis is thinking...</div>
            </div>
          </div>
        )}
        <div ref={messagesEndRef} />
      </div>

      <div className="chat-input">
        <div className="input-container">
          <textarea
            value={inputText}
            onChange={(e) => setInputText(e.target.value)}
            onKeyPress={handleKeyPress}
            placeholder="Type your message here..."
            disabled={isLoading}
            rows={1}
          />
          <button
            onClick={sendMessage}
            disabled={!inputText.trim() || isLoading}
            className="send-button"
          >
            {isLoading ? '...' : 'Send'}
          </button>
        </div>
      </div>
    </div>
  )
}

export default App
