import Link from "next/link"
import Image from "next/image"
import { css } from "styled-system/css"
import * as Avatar from "@/components/ui/avatar"
import { ModeToggle } from "../molecules/mode-toggle"

export function Header() {
  return (
    <header
      className={css({
        display: "flex",
        alignItems: "center",
        justifyContent: "space-between",
        px: "16px",
        py: "16px",
        borderBottomWidth: "1px",
        borderBottomStyle: "solid",
        borderBottomColor: "gray.4",
        color: "gray.12"
      })}
    >
      <div
        className={css({
          display: "flex",
          alignItems: "center",
          gap: "24px"
        })}
      >
        <Link
          href="/"
          className={css({
            display: "flex",
            gap: "8px",
            fontWeight: "medium"
          })}
        >
          <Image
            src="/kstm.svg"
            alt="Kstm Logo"
            width={50}
            height={24}
            priority
          />
          <span
            className={css({
              py: "8px"
            })}
          >
            kstm
          </span>
        </Link>
        <nav
          className={css({
            display: "flex",
            alignItems: "center",
            gap: "16px",
            fontSize: "sm",
            color: "gray.11"
          })}
        >
          <Link href="/">Home</Link>
          <Link href="/request-form">Request Form</Link>
        </nav>
      </div>
      <div
        className={css({
          display: "flex",
          alignItems: "center",
          gap: "8px"
        })}
      >
        <ModeToggle />
        <Avatar.Root>
          <Avatar.Fallback name="K T" />
        </Avatar.Root>
      </div>
    </header>
  )
}
