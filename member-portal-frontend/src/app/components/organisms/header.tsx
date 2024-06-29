import Link from "next/link"
import { NavigationMenuDemo } from "../molecules/navigation-menu"
import { ModeToggle } from "../molecules/mode-toggle"
import Image from "next/image"


export function Header() {
    return(
        <div className="py-4 px-4 flex items-center justify-between">
            <div className="gap-4 flex items-center">
                <Link href={"/"} className="font-medium flex gap-2">
                    <Image 
                        src="/kstm.svg"
                        alt="Kstm Logo"
                        width={50}
                        height={24}
                        priority
                    />
                    <div className="py-2">
                        kstm
                    </div>
                </Link>
                <div>
                    <NavigationMenuDemo />
                </div>
            </div>
            <div>
                <ModeToggle />
            </div>    
        </div>
    )
}