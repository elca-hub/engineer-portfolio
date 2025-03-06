import Link from "next/link";

type DPLinkProps = {
  children: React.ReactNode;
  href: string;
};

export default function DPLink(props: DPLinkProps) {
  return (
    <Link href={props.href} className="mx-1 cursor-pointer text-secondary underline transition-colors duration-300 hover:text-primary">
      {props.children}
    </Link>
  )
}
