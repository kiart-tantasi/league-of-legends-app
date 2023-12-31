import './App.css'
import React, { useEffect, useState } from 'react'

function App() {
  const [isLoading, setIsLoading] = useState(true)
  useEffect(() => {
    // simulating long loading
    setTimeout(() => {
      ;(async () => {
        try {
          const response = await fetch(
            `${process.env.REACT_APP_API_DOMAIN ?? ''}/api/health`,
          )
          if (!response.ok) {
            console.log(response.status)
          }
        } catch (err) {
          console.error(err)
        } finally {
          setIsLoading(false)
        }
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
