import type { ReactNode } from "react"

interface InputContainerProps {
  children: ReactNode
  className?: string
}

export const InputContainer = ({ children, className = "" }: InputContainerProps) => {
  return <div className={`${className}`}>{children}</div>
}
