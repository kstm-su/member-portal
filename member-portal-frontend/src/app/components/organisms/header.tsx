import Link from "next/link"
import Image from "next/image"
import { NavigationMenuDemo } from "../molecules/navigation-menu"
import { ModeToggle } from "../molecules/mode-toggle"
import { DropdownMenuDemo } from "../molecules/dropdown-menu"

export function Header() {
  return (
    <div className="py-4 px-4 flex items-center justify-between">
      <div className="gap-4 flex items-center">
        {/* kstm logoの追加 */}
        <Link href={"/"} className="font-medium flex gap-2">
          <Image
            src="/kstm.svg"
            alt="Kstm Logo"
            width={50}
            height={24}
            priority
          />
          <div className="py-2">kstm</div>
        </Link>
        <div>
          {/* ナビゲーションメニューの追加 */}
          <NavigationMenuDemo />
        </div>
      </div>
      <div className="flex gap-2">
        {/*Themeの追加*/}
        <ModeToggle />
        {/* avatarの追加 */}
        <DropdownMenuDemo />
      </div>
    </div>
  )
}
