const TEMP_MAX_WIDTH = 'max-w-[600px]'

export default function Layout({ children }: { children: JSX.Element }) {
  return <div className={`w-full ${TEMP_MAX_WIDTH} h-screen`}>{children}</div>
}
