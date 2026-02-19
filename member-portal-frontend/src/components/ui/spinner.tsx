import { ark } from '@ark-ui/react/factory'
import type { ComponentProps } from 'react'
import { styled } from 'styled-system/jsx'
import { spinner } from 'styled-system/recipes'

export const Spinner = styled(ark.span, spinner)
export type SpinnerProps = ComponentProps<typeof Spinner>
