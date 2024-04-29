import { Button } from "@/components/ui/button";
import { Drawer, DrawerContent, DrawerDescription, DrawerHeader, DrawerTitle, DrawerTrigger } from "@/components/ui/drawer";
import { Label } from "@/components/ui/label";
import { Select, SelectTrigger, SelectValue } from "@/components/ui/select";
import { Tooltip, TooltipContent, TooltipTrigger } from "@/components/ui/tooltip";
import { Triangle, SquareTerminal, Bot, Code2, Book, Settings, HelpCircle, SquareUser, Newspaper, MessageCircleMore, Radius } from "lucide-react";
import { FaDeploydog } from "react-icons/fa";

export default function Sample() {
    return <div className="text-white grid h-screen w-full pl-[56px]">
        <aside className="inset-y fixed left-0 z-20 flex h-full flex-col border-r">
            {/* home */}
            <div className="border-b p-2">
                <Button variant="outline" size="icon" aria-label="Home">
                    <Triangle className="size-5 fill-foreground" />
                </Button>
            </div>

            <nav className="grid gap-1 p-2">
                <Tooltip>
                    <TooltipTrigger asChild>
                        <Button variant="ghost" size="icon" className="rounded-lg bg-muted" aria-label="Message">
                            <MessageCircleMore className="size-5" />
                        </Button>
                    </TooltipTrigger>
                    <TooltipContent side="right" sideOffset={5}>
                        Message
                    </TooltipContent>
                </Tooltip>
                <Tooltip>
                    <TooltipTrigger asChild>
                        <Button variant="ghost" size="icon" className="rounded-lg" aria-label="Articles">
                            <Newspaper className="size-5" />
                        </Button>
                    </TooltipTrigger>
                    <TooltipContent side="right" sideOffset={5}>
                        Articles
                    </TooltipContent>
                </Tooltip>
                <Tooltip>
                    <TooltipTrigger asChild>
                        <Button variant="ghost" size="icon" className="rounded-lg" aria-label="API">
                            <Code2 className="size-5" />
                        </Button>
                    </TooltipTrigger>
                    <TooltipContent side="right" sideOffset={5}>
                        API
                    </TooltipContent>
                </Tooltip>
                <Tooltip>
                    <TooltipTrigger asChild>
                        <Button variant="ghost" size="icon" className="rounded-lg" aria-label="CI/CD">
                            <Radius className="size-5" />
                        </Button>
                    </TooltipTrigger>
                    <TooltipContent side="right" sideOffset={5}>
                        CI/CD
                    </TooltipContent>
                </Tooltip>
                <Tooltip>
                    <TooltipTrigger asChild>
                        <Button variant="ghost" size="icon" className="rounded-lg" aria-label="Documentation">
                            <Book className="size-5" />
                        </Button>
                    </TooltipTrigger>
                    <TooltipContent side="right" sideOffset={5}>
                        Documentation
                    </TooltipContent>
                </Tooltip>
                <Tooltip>
                    <TooltipTrigger asChild>
                        <Button variant="ghost" size="icon" className="rounded-lg" aria-label="Setting">
                            <Settings className="size-5" />
                        </Button>
                    </TooltipTrigger>
                    <TooltipContent side="right" sideOffset={5}>
                        Setting
                    </TooltipContent>
                </Tooltip>
            </nav>
            <nav className="mt-auto grid gap-1 p-2">
                <Tooltip>
                    <TooltipTrigger asChild>
                        <Button variant="ghost" size="icon" className="mt-auto rounded-lg" aria-label="Help">
                            <HelpCircle className="size-5" />
                        </Button>
                    </TooltipTrigger>
                    <TooltipContent side="right" sideOffset={5}>
                        Help
                    </TooltipContent>
                </Tooltip>
                <Tooltip>
                    <TooltipTrigger asChild>
                        <Button variant="ghost" size="icon" className="mt-auto rounded-lg" aria-label="Account">
                            <SquareUser className="size-5" />
                        </Button>
                    </TooltipTrigger>
                    <TooltipContent side="right" sideOffset={5}>
                        Account
                    </TooltipContent>
                </Tooltip>
            </nav>
        </aside>
        <div className="flex flex-col bg-green-100">
            <header className="sticky top-0 z-10 flex h-[57px] items-center gap-1 border-b bg-background px-4">
                <h1 className="text-xl font-semibold">Message</h1>
            </header>
            <main className="grid flex-1 gap-4 overflow-auto p-4 md:grid-cols-2 lg:grid-cols-3 bg-blue-300">
                
            </main>
        </div>
    </div>
}
