import * as React from "react";
import { Button } from "../ui/button";
import {
  Book,
  Code2,
  HelpCircle,
  MessageCircleMore,
  Newspaper,
  Radius,
  Settings,
  SquareUser,
  Triangle,
  Images,
  BookA,
} from "lucide-react";
import { Tooltip, TooltipContent, TooltipTrigger } from "../ui/tooltip";
import { Link } from "react-router-dom";

export interface NavigationProps {
  title?: string;
  children: string | React.JSX.Element | React.JSX.Element[];
  headerContent?: string | React.JSX.Element | React.JSX.Element[];
  sidebarPosition:
    | ""
    | "messages"
    | "articles"
    | "media"
    | "english_dictionary";
}

export function Navigation(props: NavigationProps) {
  const { title, children, headerContent, sidebarPosition } = props;

  return (
    <div className="text-white grid h-screen w-full pl-[56px]">
      <aside className="inset-y fixed left-0 z-20 flex h-full flex-col border-r">
        {/* home */}
        <div className="border-b p-2">
          <Link to={"/"}>
            <Button variant="outline" size="icon" aria-label="Home">
              <Triangle className="size-5 fill-foreground" />
            </Button>
          </Link>
        </div>

        <nav className="grid gap-1 p-2">
          <Tooltip>
            <TooltipTrigger asChild>
              <Button
                variant="ghost"
                size="icon"
                className={`rounded-lg ${sidebarPosition === "messages" ? "bg-muted" : ""}`}
                aria-label="Message"
              >
                <MessageCircleMore className="size-5" />
              </Button>
            </TooltipTrigger>
            <TooltipContent side="right" sideOffset={5}>
              Message
            </TooltipContent>
          </Tooltip>
          <Tooltip>
            <TooltipTrigger asChild>
              <Link to={"/articles"}>
                <Button
                  variant="ghost"
                  size="icon"
                  className={`rounded-lg ${sidebarPosition === "articles" ? "bg-muted" : ""}`}
                  aria-label="Articles"
                >
                  <Newspaper className="size-5" />
                </Button>
              </Link>
            </TooltipTrigger>
            <TooltipContent side="right" sideOffset={5}>
              Articles
            </TooltipContent>
          </Tooltip>
          <Tooltip>
            <TooltipTrigger asChild>
              <Link to={"/english_dictionary"}>
                <Button
                  variant={"ghost"}
                  size={"icon"}
                  className={`rounded-lg ${sidebarPosition === "english_dictionary" ? "bg-muted" : ""}`}
                  aria-label="EnglishDictionary"
                >
                  <BookA className="size-5" />
                </Button>
              </Link>
            </TooltipTrigger>
            <TooltipContent side="right" sideOffset={5}>
              EnglishDictionary
            </TooltipContent>
          </Tooltip>
          <Tooltip>
            <TooltipTrigger asChild>
              <Link to={"/media"}>
                <Button
                  variant="ghost"
                  size="icon"
                  className={`rounded-lg ${sidebarPosition === "media" ? "bg-muted" : ""}`}
                  aria-label="Media"
                >
                  <Images className="size-5" />
                </Button>
              </Link>
            </TooltipTrigger>
            <TooltipContent side="right" sideOffset={5}>
              Media
            </TooltipContent>
          </Tooltip>
          <Tooltip>
            <TooltipTrigger asChild>
              <Button
                variant="ghost"
                size="icon"
                className="rounded-lg"
                aria-label="API"
              >
                <Code2 className="size-5" />
              </Button>
            </TooltipTrigger>
            <TooltipContent side="right" sideOffset={5}>
              API
            </TooltipContent>
          </Tooltip>
          <Tooltip>
            <TooltipTrigger asChild>
              <Button
                variant="ghost"
                size="icon"
                className="rounded-lg"
                aria-label="CI/CD"
              >
                <Radius className="size-5" />
              </Button>
            </TooltipTrigger>
            <TooltipContent side="right" sideOffset={5}>
              CI/CD
            </TooltipContent>
          </Tooltip>
          <Tooltip>
            <TooltipTrigger asChild>
              <Button
                variant="ghost"
                size="icon"
                className="rounded-lg"
                aria-label="Documentation"
              >
                <Book className="size-5" />
              </Button>
            </TooltipTrigger>
            <TooltipContent side="right" sideOffset={5}>
              Documentation
            </TooltipContent>
          </Tooltip>
          <Tooltip>
            <TooltipTrigger asChild>
              <Button
                variant="ghost"
                size="icon"
                className="rounded-lg"
                aria-label="Setting"
              >
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
              <Button
                variant="ghost"
                size="icon"
                className="mt-auto rounded-lg"
                aria-label="Help"
              >
                <HelpCircle className="size-5" />
              </Button>
            </TooltipTrigger>
            <TooltipContent side="right" sideOffset={5}>
              Help
            </TooltipContent>
          </Tooltip>
          <Tooltip>
            <TooltipTrigger asChild>
              <Button
                variant="ghost"
                size="icon"
                className="mt-auto rounded-lg"
                aria-label="Account"
              >
                <SquareUser className="size-5" />
              </Button>
            </TooltipTrigger>
            <TooltipContent side="right" sideOffset={5}>
              Account
            </TooltipContent>
          </Tooltip>
        </nav>
      </aside>
      <div className="flex flex-col">
        <header className="sticky top-0 z-10 flex h-[57px] items-center gap-1 border-b bg-background px-4">
          <h1 className="text-xl font-semibold">{title}</h1>
          <div className="flex-grow">{headerContent && headerContent}</div>
        </header>
        <main className="grid flex-1 gap-4 overflow-auto p-4 md:grid-cols-1 lg:grid-cols-1">
          {children}
        </main>
      </div>
    </div>
  );
}
