import './App.css'
import React, { useEffect, useState } from 'react'

function App() {
  const [isLoading, setIsLoading] = useState(true)
  useEffect(() => {
    // simulating long loading
    setTimeout(() => {
      ;(async () => {
        const response = await fetch(
          `${process.env.REACT_APP_API_DOMAIN as string}/api/health`,
        )
        if (!response.ok) {
          console.log(response.status)
        }
        setIsLoading(false)
      })()
    }, 2000)
  }, [])

  return (
    <div
      className="w-full h-screen bg-blue-100 text-center"
      data-testid="root-app"
    >
      <h1>{isLoading ? 'Loading...' : 'HelloWorld'}</h1>
    </div>
  )
}

export default App
