import { ark } from '@ark-ui/react/factory'
import type { ComponentProps } from 'react'
import { styled } from 'styled-system/jsx'
import { absoluteCenter } from 'styled-system/recipes'

export const AbsoluteCenter = styled(ark.div, absoluteCenter)
export type AbsoluteCenterProps = ComponentProps<typeof AbsoluteCenter>
