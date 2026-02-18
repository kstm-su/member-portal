import { Button } from "@/components/ui/button"
import { css } from "styled-system/css"

export default function RequestFormPage() {
  return (
    <main
      className={css({
        p: "24px",
        minH: "calc(100vh - 80px)"
      })}
    >
      <h1
        className={css({
          mb: "16px",
          fontSize: "xl",
          fontWeight: "semibold"
        })}
      >
        Request Form
      </h1>
      <Button>Submit</Button>
    </main>
  )
}
