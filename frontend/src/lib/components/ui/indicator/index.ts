import { tv, type VariantProps } from "tailwind-variants";
import Indicator from "./indicator.svelte";

const indicatorVariants = tv({
    base: "flex w-3 h-3 rounded-full",
    variants: {
        variant: {
            default: "bg-gray-500",
            online: "bg-green-500",
            offline: "bg-red-500",
            warning: "bg-yellow-500"
        }
    },
    defaultVariants: {
        variant: "online"
    },
});

type Variant = VariantProps<typeof indicatorVariants>["variant"];

interface Props {
    class?: string,
    variant: Variant
}

export {
	Indicator,
    type Props,
	indicatorVariants,
};
