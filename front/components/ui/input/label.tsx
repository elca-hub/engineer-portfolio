import { DialogTrigger, Button, Popover, Dialog, Label } from "react-aria-components"
import { RiInfoI } from "react-icons/ri"
import TextWithIcon from "../text/textWithIcon"

type LabelProps = {
  title?: string
  icon?: React.ReactNode
  popoverContent?: string
  isRequired?: boolean
}

export default function DBLabel({ ...props }: LabelProps) {
  const titleNode = props.isRequired ? (
		<span>
			{props.title}
			<span className="text-red-500">*</span>
		</span>
	) : (
		props.title
	)
  const labelNode = props.icon ? <TextWithIcon icon={props.icon}>{titleNode}</TextWithIcon> : titleNode

  const labelPopOver = props.popoverContent ? (
    <DialogTrigger>
      <div className="flex flex-row items-center gap-4">
        {labelNode}
        <Button className="outline-primary rounded-full ring-1 ring-foreground text-foreground">
          <RiInfoI></RiInfoI>
        </Button>
      </div>
      <Popover
        placement="top"
        className="bg-white ring-1 ring-primary p-2 rounded entering:animate-in entering:fade-in entering:placement-bottom:slide-in-from-top-1 entering:placement-top:slide-in-from-bottom-1 exiting:animate-out exiting:fade-out exiting:placement-bottom:slide-out-to-top-1 exiting:placement-top:slide-out-to-bottom-1 fill-mode-forwards origin-top-left"
      >
        <Dialog className="text-foreground outline-none">
          <p>{props.popoverContent}</p>
        </Dialog>
      </Popover>
    </DialogTrigger>
  ) : (
    labelNode
  )

  return (
    <Label className="text-gray-700">{labelPopOver}</Label>
  )
}