import Link from "next/link";

type DPLinkProps = {
  children: React.ReactNode;
  href: string;
};


export default function DPLink(props: DPLinkProps) {
  return (
    <Link href={props.href} className="text-secondary underline mx-1 hover:text-primary cursor-pointer transition-colors duration-300">
      {props.children}
    </Link>
  )
}
