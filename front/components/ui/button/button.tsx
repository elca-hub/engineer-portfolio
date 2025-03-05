import { ButtonHTMLAttributes, DetailedHTMLProps } from "react"

interface DBButtonProps extends DetailedHTMLProps<ButtonHTMLAttributes<HTMLButtonElement>, HTMLButtonElement> {
  colormode: 'primary' | 'secondary';
}

const DPButton = ({...props}: DBButtonProps) => {
  const convertColorMode = (colormode: 'primary' | 'secondary') => {
    return colormode === 'primary' ? 'bg-primary text-foreground' : 'bg-secondary text-white';
  }

  return (
    <button
      {...props}
      className={`
        ${convertColorMode(props.colormode)}
        px-4 py-2
        font-bold text-lg
        rounded
        hover:opacity-80 hover:scale-[0.98]
        transition-all duration-300
      `}
    >
      {props.children}
    </button>
  )
}

export default DPButton
