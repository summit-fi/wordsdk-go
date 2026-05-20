package scenario

import "github.com/summit-fi/wordsdk-go/utils/ptr"

var SimpleScenario = Scenario{
	Name:   "simple",
	Locale: "en-US",
	FileSources: []FileSource{
		{Name: "browser", PathScheme: "browser/{locale}/"},
	},
	Queries: []Query{
		{Key: "history-section-label", Value: ptr.Ptr("History")},
	},
}

var BrowserScenario = Scenario{
	Name:   "browser",
	Locale: "en-US",
	FileSources: []FileSource{
		{Name: "toolkit", PathScheme: "toolkit/{locale}/"},
		{Name: "browser", PathScheme: "browser/{locale}/"},
	},
	Queries: []Query{
		{
			Key: "browser-main-window",
			Args: map[string]any{
				"content-title": "CONTENTTITLE",
			},
			Attrs: map[string]any{
				"data-title-default":         "Nightly",
				"data-title-private":         "Nightly (Private Browsing)",
				"data-content-title-default": "CONTENTTITLE — Nightly",
				"data-content-title-private": "CONTENTTITLE — Nightly (Private Browsing)",
			},
		},
		{Key: "browser-main-window-title", Value: ptr.Ptr("Nightly")},
		{Key: "window-new-shortcut", Attrs: map[string]any{"key": "N"}},
		{Key: "tab-new-shortcut", Attrs: map[string]any{"key": "T"}},
		{Key: "location-open-shortcut", Attrs: map[string]any{"key": "L"}},
		{Key: "location-open-shortcut-alt", Attrs: map[string]any{"key": "D"}},
		{Key: "search-focus-shortcut", Attrs: map[string]any{"key": "K"}},
		{Key: "search-focus-shortcut-alt", Attrs: map[string]any{"key": "J"}},
		{Key: "downloads-shortcut", Attrs: map[string]any{"key": "Y"}},
		{Key: "addons-shortcut", Attrs: map[string]any{"key": "A"}},
		{Key: "file-open-shortcut", Attrs: map[string]any{"key": "O"}},
		{Key: "save-page-shortcut", Attrs: map[string]any{"key": "S"}},
		{Key: "print-shortcut", Attrs: map[string]any{"key": "P"}},
		{Key: "close-shortcut", Attrs: map[string]any{"key": "W"}},
		{Key: "close-shortcut", Attrs: map[string]any{"key": "W"}},
		{Key: "mute-toggle-shortcut", Attrs: map[string]any{"key": "M"}},
		{Key: "text-action-undo-shortcut", Attrs: map[string]any{"key": "Z"}},
		{Key: "text-action-undo-shortcut", Attrs: map[string]any{"key": "Z"}},
		{Key: "text-action-cut-shortcut", Attrs: map[string]any{"key": "X"}},
		{Key: "text-action-copy-shortcut", Attrs: map[string]any{"key": "C"}},
		{Key: "text-action-paste-shortcut", Attrs: map[string]any{"key": "V"}},
		{Key: "text-action-select-all-shortcut", Attrs: map[string]any{"key": "A"}},
		{Key: "nav-back-shortcut-alt", Attrs: map[string]any{"key": "["}},
		{Key: "nav-fwd-shortcut-alt", Attrs: map[string]any{"key": "]"}},
		{Key: "history-show-all-shortcut", Attrs: map[string]any{"key": "H"}},
		{Key: "reader-mode-toggle-shortcut-other", Attrs: map[string]any{"key": "R"}},
		{Key: "picture-in-picture-toggle-shortcut", Attrs: map[string]any{"key": "]"}},
		{Key: "picture-in-picture-toggle-shortcut-alt", Attrs: map[string]any{"key": "}"}},
		{Key: "nav-reload-shortcut", Attrs: map[string]any{"key": "R"}},
		{Key: "nav-reload-shortcut", Attrs: map[string]any{"key": "R"}},
		{Key: "page-source-shortcut", Attrs: map[string]any{"key": "U"}},
		{Key: "page-info-shortcut", Attrs: map[string]any{"key": "I"}},
		{Key: "find-shortcut", Attrs: map[string]any{"key": "F"}},
		{Key: "search-find-again-shortcut", Attrs: map[string]any{"key": "G"}},
		{Key: "search-find-again-shortcut", Attrs: map[string]any{"key": "G"}},
		{Key: "search-find-again-shortcut-alt", Attrs: map[string]any{"keycode": "VK_F3"}},
		{Key: "search-find-again-shortcut-alt", Attrs: map[string]any{"keycode": "VK_F3"}},
		{Key: "bookmark-this-page-shortcut", Attrs: map[string]any{"key": "D"}},
		{Key: "bookmark-this-page-shortcut", Attrs: map[string]any{"key": "D"}},
		{Key: "bookmark-show-library-shortcut", Attrs: map[string]any{"key": "O"}},
		{Key: "bookmark-show-sidebar-shortcut", Attrs: map[string]any{"key": "B"}},
		{Key: "bookmark-show-toolbar-shortcut", Attrs: map[string]any{"key": "B"}},
		{Key: "history-sidebar-shortcut", Attrs: map[string]any{"key": "H"}},
		{Key: "full-zoom-reduce-shortcut", Attrs: map[string]any{"key": "-"}},
		{Key: "full-zoom-reduce-shortcut-alt-a", Attrs: map[string]any{"key": "_"}},
		{Key: "full-zoom-reduce-shortcut-alt-b", Attrs: map[string]any{"key": ""}},
		{Key: "full-zoom-enlarge-shortcut", Attrs: map[string]any{"key": "+"}},
		{Key: "full-zoom-enlarge-shortcut-alt", Attrs: map[string]any{"key": "="}},
		{Key: "full-zoom-enlarge-shortcut-alt2", Attrs: map[string]any{"key": ""}},
		{Key: "full-zoom-reset-shortcut", Attrs: map[string]any{"key": "0"}},
		{Key: "full-zoom-reset-shortcut-alt", Attrs: map[string]any{"key": ""}},
		{Key: "bidi-switch-direction-shortcut", Attrs: map[string]any{"key": "X"}},
		{Key: "private-browsing-shortcut", Attrs: map[string]any{"key": "P"}},
		{Key: "quit-app-shortcut", Attrs: map[string]any{"key": "Q"}},
		{Key: "tab-new-shortcut", Attrs: map[string]any{"key": "T"}},
		{Key: "window-new-shortcut", Attrs: map[string]any{"key": "N"}},
		{Key: "sidebar-menu-bookmarks", Attrs: map[string]any{"label": "Bookmarks"}},
		{Key: "sidebar-menu-history", Attrs: map[string]any{"label": "History"}},
		{Key: "sidebar-menu-synced-tabs", Attrs: map[string]any{"label": "Synced Tabs"}},
		{Key: "sidebar-menu-close", Attrs: map[string]any{"label": "Close Sidebar"}},
		{Key: "full-screen-autohide", Attrs: map[string]any{"label": "Hide Toolbars", "accesskey": "H"}},
		{
			Key: "full-screen-exit",
			Attrs: map[string]any{
				"label":     "Exit Full Screen Mode",
				"accesskey": "F",
			},
		},
		{
			Key: "main-context-menu-back",
			Attrs: map[string]any{
				"tooltiptext": "Go back one page",
				"aria-label":  "Back",
				"accesskey":   "B",
			},
		},
		{
			Key: "main-context-menu-forward",
			Attrs: map[string]any{
				"tooltiptext": "Go forward one page",
				"aria-label":  "Forward",
				"accesskey":   "F",
			},
		},
		{
			Key: "main-context-menu-reload",
			Attrs: map[string]any{
				"aria-label": "Reload",
				"accesskey":  "R",
			},
		},
		{
			Key: "main-context-menu-stop",
			Attrs: map[string]any{
				"aria-label": "Stop",
				"accesskey":  "S",
			},
		},
		{
			Key: "main-context-menu-bookmark-add",
			Attrs: map[string]any{
				"aria-label":  "Bookmark This Page",
				"accesskey":   "m",
				"tooltiptext": "Bookmark this page",
			},
		},
		{
			Key: "main-context-menu-open-link",
			Attrs: map[string]any{
				"label":     "Open Link",
				"accesskey": "O",
			},
		},
		{
			Key: "main-context-menu-open-link-new-tab",
			Attrs: map[string]any{
				"label":     "Open Link in New Tab",
				"accesskey": "T",
			},
		},
		{
			Key: "main-context-menu-open-link-container-tab",
			Attrs: map[string]any{
				"label":     "Open Link in New Container Tab",
				"accesskey": "b",
			},
		},
		{
			Key: "main-context-menu-open-link-new-window",
			Attrs: map[string]any{
				"label":     "Open Link in New Window",
				"accesskey": "W",
			},
		},
		{
			Key: "main-context-menu-open-link-new-private-window",
			Attrs: map[string]any{
				"label":     "Open Link in New Private Window",
				"accesskey": "P",
			},
		},
		{
			Key: "main-context-menu-bookmark-this-link",
			Attrs: map[string]any{
				"label":     "Bookmark This Link",
				"accesskey": "L",
			},
		},
		{
			Key: "main-context-menu-save-link",
			Attrs: map[string]any{
				"label":     "Save Link As…",
				"accesskey": "k",
			},
		},
		{
			Key: "main-context-menu-save-link-to-pocket",
			Attrs: map[string]any{
				"label":     "Save Link to Pocket",
				"accesskey": "o",
			},
		},
		{
			Key: "main-context-menu-copy-email",
			Attrs: map[string]any{
				"label":     "Copy Email Address",
				"accesskey": "A",
			},
		},
		{
			Key: "main-context-menu-copy-link",
			Attrs: map[string]any{
				"label":     "Copy Link Location",
				"accesskey": "a",
			},
		},
		{
			Key: "main-context-menu-media-play",
			Attrs: map[string]any{
				"label":     "Play",
				"accesskey": "P",
			},
		},
		{
			Key: "main-context-menu-media-pause",
			Attrs: map[string]any{
				"label":     "Pause",
				"accesskey": "P",
			},
		},
		{
			Key: "main-context-menu-media-mute",
			Attrs: map[string]any{
				"label":     "Mute",
				"accesskey": "M",
			},
		},
		{
			Key: "main-context-menu-media-unmute",
			Attrs: map[string]any{
				"label":     "Unmute",
				"accesskey": "m",
			},
		},
		{
			Key: "main-context-menu-media-play-speed",
			Attrs: map[string]any{
				"label":     "Play Speed",
				"accesskey": "d",
			},
		},
		{
			Key: "main-context-menu-media-play-speed-slow",
			Attrs: map[string]any{
				"label":     "Slow (0.5)",
				"accesskey": "S",
			},
		},
		{
			Key: "main-context-menu-media-play-speed-normal",
			Attrs: map[string]any{
				"label":     "Normal",
				"accesskey": "N",
			},
		},
		{
			Key: "main-context-menu-media-play-speed-fast",
			Attrs: map[string]any{
				"label":     "Fast (1.25)",
				"accesskey": "F",
			},
		},
		{
			Key: "main-context-menu-media-play-speed-faster",
			Attrs: map[string]any{
				"label":     "Faster (1.5)",
				"accesskey": "a",
			},
		},
		{
			Key: "main-context-menu-media-play-speed-fastest",
			Attrs: map[string]any{
				"label":     "Ludicrous (2)",
				"accesskey": "L",
			},
		},
		{
			Key: "main-context-menu-media-loop",
			Attrs: map[string]any{
				"label":     "Loop",
				"accesskey": "L",
			},
		},
		{
			Key: "main-context-menu-media-show-controls",
			Attrs: map[string]any{
				"label":     "Show Controls",
				"accesskey": "C",
			},
		},
		{
			Key: "main-context-menu-media-hide-controls",
			Attrs: map[string]any{
				"label":     "Hide Controls",
				"accesskey": "C",
			},
		},
		{
			Key: "main-context-menu-media-video-fullscreen",
			Attrs: map[string]any{
				"label":     "Full Screen",
				"accesskey": "F",
			},
		},
		{
			Key: "main-context-menu-media-video-leave-fullscreen",
			Attrs: map[string]any{
				"label":     "Exit Full Screen",
				"accesskey": "u",
			},
		},
		{
			Key: "main-context-menu-media-pip",
			Attrs: map[string]any{
				"label":     "Picture-in-Picture",
				"accesskey": "u",
			},
		},
		{
			Key: "main-context-menu-image-reload",
			Attrs: map[string]any{
				"label":     "Reload Image",
				"accesskey": "R",
			},
		},
		{
			Key: "main-context-menu-image-view",
			Attrs: map[string]any{
				"label":     "View Image",
				"accesskey": "I",
			},
		},
		{
			Key: "main-context-menu-video-view",
			Attrs: map[string]any{
				"label":     "View Video",
				"accesskey": "i",
			},
		},
		{
			Key: "main-context-menu-image-copy",
			Attrs: map[string]any{
				"label":     "Copy Image",
				"accesskey": "y",
			},
		},
		{
			Key: "main-context-menu-image-copy-location",
			Attrs: map[string]any{
				"label":     "Copy Image Location",
				"accesskey": "o",
			},
		},
		{
			Key: "main-context-menu-video-copy-location",
			Attrs: map[string]any{
				"label":     "Copy Video Location",
				"accesskey": "o",
			},
		},
		{
			Key: "main-context-menu-audio-copy-location",
			Attrs: map[string]any{
				"label":     "Copy Audio Location",
				"accesskey": "o",
			},
		},
		{
			Key: "main-context-menu-image-save-as",
			Attrs: map[string]any{
				"label":     "Save Image As…",
				"accesskey": "v",
			},
		},
		{
			Key: "main-context-menu-image-email",
			Attrs: map[string]any{
				"label":     "Email Image…",
				"accesskey": "g",
			},
		},
		{
			Key: "main-context-menu-image-set-as-background",
			Attrs: map[string]any{
				"label":     "Set As Desktop Background…",
				"accesskey": "S",
			},
		},
		{
			Key: "main-context-menu-image-info",
			Attrs: map[string]any{
				"label":     "View Image Info",
				"accesskey": "f",
			},
		},
		{
			Key: "main-context-menu-image-desc",
			Attrs: map[string]any{
				"label":     "View Description",
				"accesskey": "D",
			},
		},
		{
			Key: "main-context-menu-video-save-as",
			Attrs: map[string]any{
				"label":     "Save Video As…",
				"accesskey": "v",
			},
		},
		{
			Key: "main-context-menu-audio-save-as",
			Attrs: map[string]any{
				"label":     "Save Audio As…",
				"accesskey": "v",
			},
		},
		{
			Key: "main-context-menu-video-image-save-as",
			Attrs: map[string]any{
				"label":     "Save Snapshot As…",
				"accesskey": "S",
			},
		},
		{
			Key: "main-context-menu-video-email",
			Attrs: map[string]any{
				"label":     "Email Video…",
				"accesskey": "a",
			},
		},
		{
			Key: "main-context-menu-audio-email",
			Attrs: map[string]any{
				"label":     "Email Audio…",
				"accesskey": "a",
			},
		},
		{
			Key: "main-context-menu-plugin-play",
			Attrs: map[string]any{
				"label":     "Activate this plugin",
				"accesskey": "c",
			},
		},
		{
			Key: "main-context-menu-plugin-hide",
			Attrs: map[string]any{
				"label":     "Hide this plugin",
				"accesskey": "H",
			},
		},
		{
			Key: "main-context-menu-page-save",
			Attrs: map[string]any{
				"label":     "Save Page As…",
				"accesskey": "P",
			},
		},
		{
			Key: "main-context-menu-save-to-pocket",
			Attrs: map[string]any{
				"label":     "Save Page to Pocket",
				"accesskey": "k",
			},
		},
		{
			Key: "main-context-menu-send-to-device",
			Attrs: map[string]any{
				"label":     "Send Page to Device",
				"accesskey": "n",
			},
		},
		{
			Key: "main-context-menu-view-background-image",
			Attrs: map[string]any{
				"label":     "View Background Image",
				"accesskey": "w",
			},
		},
		{
			Key: "main-context-menu-generate-new-password",
			Attrs: map[string]any{
				"label":     "Use Generated Password…",
				"accesskey": "G",
			},
		},
		{
			Key: "text-action-undo",
			Attrs: map[string]any{
				"label":     "Undo",
				"accesskey": "U",
			},
		},
		{
			Key: "text-action-cut",
			Attrs: map[string]any{
				"label":     "Cut",
				"accesskey": "t",
			},
		},
		{
			Key: "text-action-copy",
			Attrs: map[string]any{
				"label":     "Copy",
				"accesskey": "C",
			},
		},
		{
			Key: "text-action-paste",
			Attrs: map[string]any{
				"label":     "Paste",
				"accesskey": "P",
			},
		},
		{
			Key: "text-action-delete",
			Attrs: map[string]any{
				"label":     "Delete",
				"accesskey": "D",
			},
		},
		{
			Key: "text-action-select-all",
			Attrs: map[string]any{
				"label":     "Select All",
				"accesskey": "A",
			},
		},
		{
			Key: "main-context-menu-keyword",
			Attrs: map[string]any{
				"label":     "Add a Keyword for this Search…",
				"accesskey": "K",
			},
		},
		{
			Key: "main-context-menu-link-send-to-device",
			Attrs: map[string]any{
				"label":     "Send Link to Device",
				"accesskey": "n",
			},
		},
		{
			Key: "main-context-menu-frame",
			Attrs: map[string]any{
				"label":     "This Frame",
				"accesskey": "h",
			},
		},
		{
			Key: "main-context-menu-frame-show-this",
			Attrs: map[string]any{
				"label":     "Show Only This Frame",
				"accesskey": "S",
			},
		},
		{
			Key: "main-context-menu-frame-open-tab",
			Attrs: map[string]any{
				"label":     "Open Frame in New Tab",
				"accesskey": "T",
			},
		},
		{
			Key: "main-context-menu-frame-open-window",
			Attrs: map[string]any{
				"label":     "Open Frame in New Window",
				"accesskey": "W",
			},
		},
		{
			Key: "main-context-menu-frame-reload",
			Attrs: map[string]any{
				"label":     "Reload Frame",
				"accesskey": "R",
			},
		},
		{
			Key: "main-context-menu-frame-bookmark",
			Attrs: map[string]any{
				"label":     "Bookmark This Frame",
				"accesskey": "m",
			},
		},
		{
			Key: "main-context-menu-frame-save-as",
			Attrs: map[string]any{
				"label":     "Save Frame As…",
				"accesskey": "F",
			},
		},
		{
			Key: "main-context-menu-frame-print",
			Attrs: map[string]any{
				"label":     "Print Frame…",
				"accesskey": "P",
			},
		},
		{
			Key: "main-context-menu-frame-view-source",
			Attrs: map[string]any{
				"label":     "View Frame Source",
				"accesskey": "V",
			},
		},
		{
			Key: "main-context-menu-frame-view-info",
			Attrs: map[string]any{
				"label":     "View Frame Info",
				"accesskey": "I",
			},
		},
		{
			Key: "main-context-menu-print-selection",
			Attrs: map[string]any{
				"label":     "Print Selection",
				"accesskey": "r",
			},
		},
		{
			Key: "main-context-menu-view-selection-source",
			Attrs: map[string]any{
				"label":     "View Selection Source",
				"accesskey": "e",
			},
		},
		{
			Key: "main-context-menu-view-page-source",
			Attrs: map[string]any{
				"label":     "View Page Source",
				"accesskey": "V",
			},
		},
		{
			Key: "main-context-menu-view-page-info",
			Attrs: map[string]any{
				"label":     "View Page Info",
				"accesskey": "I",
			},
		},
		{
			Key: "main-context-menu-bidi-switch-text",
			Attrs: map[string]any{
				"label":     "Switch Text Direction",
				"accesskey": "w",
			},
		},
		{
			Key: "main-context-menu-bidi-switch-page",
			Attrs: map[string]any{
				"label":     "Switch Page Direction",
				"accesskey": "D",
			},
		},
		{Key: "main-context-menu-inspect-a11y-properties", Attrs: map[string]any{"label": "Inspect Accessibility Properties"}},
		{
			Key: "main-context-menu-inspect-element",
			Attrs: map[string]any{
				"label":     "Inspect Element",
				"accesskey": "Q",
			},
		},
		{
			Key: "main-context-menu-eme-learn-more",
			Attrs: map[string]any{
				"label":     "Learn more about DRM…",
				"accesskey": "D",
			},
		},
		{
			Key: "places-open",
			Attrs: map[string]any{
				"label":     "Open",
				"accesskey": "O",
			},
		},
		{
			Key: "places-open-tab",
			Attrs: map[string]any{
				"label":     "Open in a New Tab",
				"accesskey": "w",
			},
		},
		{
			Key: "places-open-all-in-tabs",
			Attrs: map[string]any{
				"label":     "Open All in Tabs",
				"accesskey": "O",
			},
		},
		{
			Key: "places-open-all-in-tabs",
			Attrs: map[string]any{
				"label":     "Open All in Tabs",
				"accesskey": "O",
			},
		},
		{
			Key: "places-open-window",
			Attrs: map[string]any{
				"label":     "Open in a New Window",
				"accesskey": "N",
			},
		},
		{
			Key: "places-open-private-window",
			Attrs: map[string]any{
				"label":     "Open in a New Private Window",
				"accesskey": "P",
			},
		},
		{
			Key: "places-new-bookmark",
			Attrs: map[string]any{
				"label":     "New Bookmark…",
				"accesskey": "B",
			},
		},
		{
			Key: "places-new-folder-contextmenu",
			Attrs: map[string]any{
				"label":     "New Folder…",
				"accesskey": "F",
			},
		},
		{
			Key: "places-new-separator",
			Attrs: map[string]any{
				"label":     "New Separator",
				"accesskey": "S",
			},
		},
		{
			Key: "text-action-cut",
			Attrs: map[string]any{
				"label":     "Cut",
				"accesskey": "t",
			},
		},
		{
			Key: "text-action-copy",
			Attrs: map[string]any{
				"label":     "Copy",
				"accesskey": "C",
			},
		},
		{
			Key: "text-action-paste",
			Attrs: map[string]any{
				"label":     "Paste",
				"accesskey": "P",
			},
		},
		{
			Key: "text-action-delete",
			Attrs: map[string]any{
				"label":     "Delete",
				"accesskey": "D",
			},
		},
		{
			Key: "places-delete-domain-data",
			Attrs: map[string]any{
				"label":     "Forget About This Site",
				"accesskey": "F",
			},
		},
		{
			Key: "places-sortby-name",
			Attrs: map[string]any{
				"label":     "Sort By Name",
				"accesskey": "r",
			},
		},
		{
			Key: "places-properties",
			Attrs: map[string]any{
				"label":     "Properties",
				"accesskey": "i",
			},
		},
		{Key: "page-action-add-to-urlbar", Attrs: map[string]any{"label": "Add to Address Bar"}},
		{Key: "page-action-remove-from-urlbar", Attrs: map[string]any{"label": "Remove from Address Bar"}},
		{Key: "page-action-add-to-urlbar", Attrs: map[string]any{"label": "Add to Address Bar"}},
		{Key: "page-action-remove-from-urlbar", Attrs: map[string]any{"label": "Remove from Address Bar"}},
		{Key: "page-action-manage-extension", Attrs: map[string]any{"label": "Manage Extension…"}},
		{Key: "page-action-remove-extension", Attrs: map[string]any{"label": "Remove Extension"}},
		{Key: "navbar-tooltip-back", Attrs: map[string]any{"value": "Go back one page"}},
		{Key: "navbar-tooltip-instruction", Attrs: map[string]any{"value": "Right-click or pull down to show history"}},
		{Key: "navbar-tooltip-forward", Attrs: map[string]any{"value": "Go forward one page"}},
		{Key: "navbar-tooltip-instruction", Attrs: map[string]any{"value": "Right-click or pull down to show history"}},
		{
			Key: "popup-select-camera",
			Attrs: map[string]any{
				"value":     "Camera to share:",
				"accesskey": "C",
			},
		},
		{Key: "popup-all-windows-shared", Value: ptr.Ptr("All visible windows on your screen will be shared.")},
		{
			Key: "popup-select-microphone",
			Attrs: map[string]any{
				"value":     "Microphone to share:",
				"accesskey": "M",
			},
		},
		{Key: "downloads-panel", Attrs: map[string]any{"aria-label": "Downloads"}},
		{
			Key: "downloads-cmd-pause",
			Attrs: map[string]any{
				"label":     "Pause",
				"accesskey": "P",
			},
		},
		{
			Key: "downloads-cmd-resume",
			Attrs: map[string]any{
				"label":     "Resume",
				"accesskey": "R",
			},
		},
		{
			Key: "downloads-cmd-unblock",
			Attrs: map[string]any{
				"label":     "Allow Download",
				"accesskey": "o",
			},
		},
		{
			Key: "downloads-cmd-use-system-default",
			Attrs: map[string]any{
				"label":     "Open In System Viewer",
				"accesskey": "V",
			},
		},
		{
			Key: "downloads-cmd-always-use-system-default",
			Attrs: map[string]any{
				"label":     "Always Open In System Viewer",
				"accesskey": "w",
			},
		},
		{
			Key: "downloads-cmd-show-menuitem",
			Attrs: map[string]any{
				"label":     "Open Containing Folder",
				"accesskey": "F",
			},
		},
		{
			Key: "downloads-cmd-go-to-download-page",
			Attrs: map[string]any{
				"label":     "Go To Download Page",
				"accesskey": "G",
			},
		},
		{
			Key: "downloads-cmd-copy-download-link",
			Attrs: map[string]any{
				"label":     "Copy Download Link",
				"accesskey": "L",
			},
		},
		{
			Key: "downloads-cmd-remove-from-history",
			Attrs: map[string]any{
				"label":     "Remove From History",
				"accesskey": "e",
			},
		},
		{
			Key: "downloads-cmd-clear-list",
			Attrs: map[string]any{
				"label":     "Clear Preview Panel",
				"accesskey": "a",
			},
		},
		{
			Key: "downloads-cmd-clear-downloads",
			Attrs: map[string]any{
				"label":     "Clear Downloads",
				"accesskey": "D",
			},
		},
		{Key: "downloads-panel-list", Attrs: map[string]any{"style": "width: 70ch"}},
		{Key: "downloads-panel-empty", Attrs: map[string]any{"value": "No downloads for this session."}},
		{
			Key: "downloads-history",
			Attrs: map[string]any{
				"label":     "Show All Downloads",
				"accesskey": "S",
			},
		},
		{Key: "downloads-details", Attrs: map[string]any{"title": "Download Details"}},
		{Key: "downloads-cmd-show-downloads", Attrs: map[string]any{"label": "Show Downloads Folder"}},
		{
			Key: "downloads-history",
			Attrs: map[string]any{
				"label":     "Show All Downloads",
				"accesskey": "S",
			},
		},
		{
			Key: "text-action-undo",
			Attrs: map[string]any{
				"label":     "Undo",
				"accesskey": "U",
			},
		},
		{
			Key: "text-action-cut",
			Attrs: map[string]any{
				"label":     "Cut",
				"accesskey": "t",
			},
		},
		{
			Key: "text-action-copy",
			Attrs: map[string]any{
				"label":     "Copy",
				"accesskey": "C",
			},
		},
		{
			Key: "text-action-paste",
			Attrs: map[string]any{
				"label":     "Paste",
				"accesskey": "P",
			},
		},
		{
			Key: "text-action-delete",
			Attrs: map[string]any{
				"label":     "Delete",
				"accesskey": "D",
			},
		},
		{
			Key: "text-action-select-all",
			Attrs: map[string]any{
				"label":     "Select All",
				"accesskey": "A",
			},
		},
		{
			Key: "menu-file",
			Attrs: map[string]any{
				"label":     "File",
				"accesskey": "F",
			},
		},
		{
			Key: "menu-file-new-tab",
			Attrs: map[string]any{
				"label":     "New Tab",
				"accesskey": "T",
			},
		},
		{
			Key: "menu-file-new-container-tab",
			Attrs: map[string]any{
				"label":     "New Container Tab",
				"accesskey": "b",
			},
		},
		{
			Key: "menu-file-new-window",
			Attrs: map[string]any{
				"label":     "New Window",
				"accesskey": "N",
			},
		},
		{
			Key: "menu-file-new-private-window",
			Attrs: map[string]any{
				"label":     "New Private Window",
				"accesskey": "W",
			},
		},
		{Key: "menu-file-open-location", Attrs: map[string]any{"label": "Open Location…"}},
		{
			Key: "menu-file-open-file",
			Attrs: map[string]any{
				"label":     "Open File…",
				"accesskey": "O",
			},
		},
		{
			Key: "menu-file-close",
			Attrs: map[string]any{
				"label":     "Close",
				"accesskey": "C",
			},
		},
		{
			Key: "menu-file-close-window",
			Attrs: map[string]any{
				"label":     "Close Window",
				"accesskey": "d",
			},
		},
		{
			Key: "menu-file-save-page",
			Attrs: map[string]any{
				"label":     "Save Page As…",
				"accesskey": "A",
			},
		},
		{
			Key: "menu-file-email-link",
			Attrs: map[string]any{
				"label":     "Email Link…",
				"accesskey": "E",
			},
		},
		{
			Key: "menu-file-print-preview",
			Attrs: map[string]any{
				"label":     "Print Preview",
				"accesskey": "v",
			},
		},
		{
			Key: "menu-file-print",
			Attrs: map[string]any{
				"label":     "Print…",
				"accesskey": "P",
			},
		},
		{
			Key: "menu-file-import-from-another-browser",
			Attrs: map[string]any{
				"label":     "Import From Another Browser…",
				"accesskey": "I",
			},
		},
		{
			Key: "menu-file-go-offline",
			Attrs: map[string]any{
				"label":     "Work Offline",
				"accesskey": "k",
			},
		},
		{
			Key: "menu-edit",
			Attrs: map[string]any{
				"label":     "Edit",
				"accesskey": "E",
			},
		},
		{
			Key: "text-action-undo",
			Attrs: map[string]any{
				"label":     "Undo",
				"accesskey": "U",
			},
		},
		{
			Key: "text-action-redo",
			Attrs: map[string]any{
				"label":     "Redo",
				"accesskey": "R",
			},
		},
		{
			Key: "text-action-cut",
			Attrs: map[string]any{
				"label":     "Cut",
				"accesskey": "t",
			},
		},
		{
			Key: "text-action-copy",
			Attrs: map[string]any{
				"label":     "Copy",
				"accesskey": "C",
			},
		},
		{
			Key: "text-action-paste",
			Attrs: map[string]any{
				"label":     "Paste",
				"accesskey": "P",
			},
		},
		{
			Key: "text-action-delete",
			Attrs: map[string]any{
				"label":     "Delete",
				"accesskey": "D",
			},
		},
		{
			Key: "text-action-select-all",
			Attrs: map[string]any{
				"label":     "Select All",
				"accesskey": "A",
			},
		},
		{
			Key: "menu-edit-find-on",
			Attrs: map[string]any{
				"label":     "Find in This Page…",
				"accesskey": "F",
			},
		},
		{
			Key: "menu-edit-find-again",
			Attrs: map[string]any{
				"label":     "Find Again",
				"accesskey": "g",
			},
		},
		{
			Key: "menu-edit-bidi-switch-text-direction",
			Attrs: map[string]any{
				"label":     "Switch Text Direction",
				"accesskey": "w",
			},
		},
		{
			Key: "menu-preferences",
			Attrs: map[string]any{
				"label":     "Preferences",
				"accesskey": "n",
			},
		},
		{
			Key: "menu-view",
			Attrs: map[string]any{
				"label":     "View",
				"accesskey": "V",
			},
		},
		{
			Key: "menu-view-toolbars-menu",
			Attrs: map[string]any{
				"label":     "Toolbars",
				"accesskey": "T",
			},
		},
		{
			Key: "menu-view-customize-toolbar",
			Attrs: map[string]any{
				"label":     "Customize…",
				"accesskey": "C",
			},
		},
		{
			Key: "menu-view-sidebar",
			Attrs: map[string]any{
				"label":     "Sidebar",
				"accesskey": "e",
			},
		},
		{Key: "menu-view-bookmarks", Attrs: map[string]any{"label": "Bookmarks"}},
		{Key: "menu-view-history-button", Attrs: map[string]any{"label": "History"}},
		{Key: "menu-view-synced-tabs-sidebar", Attrs: map[string]any{"label": "Synced Tabs"}},
		{
			Key: "menu-view-full-zoom",
			Attrs: map[string]any{
				"label":     "Zoom",
				"accesskey": "Z",
			},
		},
		{
			Key: "menu-view-full-zoom-enlarge",
			Attrs: map[string]any{
				"label":     "Zoom In",
				"accesskey": "I",
			},
		},
		{
			Key: "menu-view-full-zoom-reduce",
			Attrs: map[string]any{
				"label":     "Zoom Out",
				"accesskey": "O",
			},
		},
		{
			Key: "menu-view-full-zoom-actual-size",
			Attrs: map[string]any{
				"label":     "Actual Size",
				"accesskey": "A",
			},
		},
		{
			Key: "menu-view-full-zoom-toggle",
			Attrs: map[string]any{
				"label":     "Zoom Text Only",
				"accesskey": "T",
			},
		},
		{
			Key: "menu-view-page-style-menu",
			Attrs: map[string]any{
				"label":     "Page Style",
				"accesskey": "y",
			},
		},
		{
			Key: "menu-view-page-style-no-style",
			Attrs: map[string]any{
				"label":     "No Style",
				"accesskey": "n",
			},
		},
		{
			Key: "menu-view-page-basic-style",
			Attrs: map[string]any{
				"label":     "Basic Page Style",
				"accesskey": "B",
			},
		},
		{
			Key: "menu-view-charset",
			Attrs: map[string]any{
				"label":     "Text Encoding",
				"accesskey": "c",
			},
		},
		{
			Key: "menu-view-full-screen",
			Attrs: map[string]any{
				"label":     "Full Screen",
				"accesskey": "F",
			},
		},
		{
			Key: "menu-view-show-all-tabs",
			Attrs: map[string]any{
				"label":     "Show All Tabs",
				"accesskey": "A",
			},
		},
		{
			Key: "menu-view-bidi-switch-page-direction",
			Attrs: map[string]any{
				"label":     "Switch Page Direction",
				"accesskey": "D",
			},
		},
		{
			Key: "menu-history",
			Attrs: map[string]any{
				"label":     "History",
				"accesskey": "s",
			},
		},
		{Key: "menu-history-show-all-history", Attrs: map[string]any{"label": "Show All History"}},
		{Key: "menu-history-clear-recent-history", Attrs: map[string]any{"label": "Clear Recent History…"}},
		{Key: "menu-history-synced-tabs", Attrs: map[string]any{"label": "Synced Tabs"}},
		{Key: "menu-history-restore-last-session", Attrs: map[string]any{"label": "Restore Previous Session"}},
		{Key: "menu-history-hidden-tabs", Attrs: map[string]any{"label": "Hidden Tabs"}},
		{Key: "menu-history-undo-menu", Attrs: map[string]any{"label": "Recently Closed Tabs"}},
		{Key: "menu-history-undo-window-menu", Attrs: map[string]any{"label": "Recently Closed Windows"}},
		{
			Key: "menu-bookmarks-menu",
			Attrs: map[string]any{
				"label":     "Bookmarks",
				"accesskey": "B",
			},
		},
		{Key: "menu-bookmarks-show-all", Attrs: map[string]any{"label": "Show All Bookmarks"}},
		{Key: "menu-bookmark-this-page", Attrs: map[string]any{"label": "Bookmark This Page"}},
		{Key: "menu-bookmarks-all-tabs", Attrs: map[string]any{"label": "Bookmark All Tabs…"}},
		{Key: "menu-bookmarks-toolbar", Attrs: map[string]any{"label": "Bookmarks Toolbar"}},
		{Key: "menu-bookmarks-other", Attrs: map[string]any{"label": "Other Bookmarks"}},
		{Key: "menu-bookmarks-mobile", Attrs: map[string]any{"label": "Mobile Bookmarks"}},
		{
			Key: "menu-tools",
			Attrs: map[string]any{
				"label":     "Tools",
				"accesskey": "T",
			},
		},
		{
			Key: "menu-tools-downloads",
			Attrs: map[string]any{
				"label":     "Downloads",
				"accesskey": "D",
			},
		},
		{
			Key: "menu-tools-addons",
			Attrs: map[string]any{
				"label":     "Add-ons",
				"accesskey": "A",
			},
		},
		{
			Key: "menu-tools-fxa-sign-in",
			Attrs: map[string]any{
				"label":     "Sign In To Firefox…",
				"accesskey": "g",
			},
		},
		{
			Key: "menu-tools-turn-on-sync",
			Attrs: map[string]any{
				"label":     "Turn on Sync…",
				"accesskey": "n",
			},
		},
		{
			Key: "menu-tools-fxa-sign-in",
			Attrs: map[string]any{
				"label":     "Sign In To Firefox…",
				"accesskey": "g",
			},
		},
		{
			Key: "menu-tools-sync-now",
			Attrs: map[string]any{
				"label":     "Sync Now",
				"accesskey": "S",
			},
		},
		{
			Key: "menu-tools-fxa-re-auth",
			Attrs: map[string]any{
				"label":     "Reconnect to Firefox…",
				"accesskey": "R",
			},
		},
		{
			Key: "menu-tools-web-developer",
			Attrs: map[string]any{
				"label":     "Web Developer",
				"accesskey": "W",
			},
		},
		{
			Key: "menu-tools-page-source",
			Attrs: map[string]any{
				"label":     "Page Source",
				"accesskey": "o",
			},
		},
		{
			Key: "menu-tools-page-info",
			Attrs: map[string]any{
				"label":     "Page Info",
				"accesskey": "I",
			},
		},
		{
			Key: "menu-help",
			Attrs: map[string]any{
				"label":     "Help",
				"accesskey": "H",
			},
		},
		{
			Key: "menu-help-product",
			Attrs: map[string]any{
				"label":     "Nightly Help",
				"accesskey": "H",
			},
		},
		{
			Key: "menu-help-show-tour",
			Attrs: map[string]any{
				"label":     "Nightly Tour",
				"accesskey": "o",
			},
		},
		{
			Key: "menu-help-import-from-another-browser",
			Attrs: map[string]any{
				"label":     "Import From Another Browser…",
				"accesskey": "I",
			},
		},
		{
			Key: "menu-help-keyboard-shortcuts",
			Attrs: map[string]any{
				"label":     "Keyboard Shortcuts",
				"accesskey": "K",
			},
		},
		{
			Key: "menu-help-troubleshooting-info",
			Attrs: map[string]any{
				"label":     "Troubleshooting Information",
				"accesskey": "T",
			},
		},
		{
			Key: "menu-help-feedback-page",
			Attrs: map[string]any{
				"label":     "Submit Feedback…",
				"accesskey": "S",
			},
		},
		{
			Key: "menu-help-safe-mode-without-addons",
			Attrs: map[string]any{
				"label":     "Restart With Add-ons Disabled…",
				"accesskey": "R",
			},
		},
		{
			Key: "menu-help-report-deceptive-site",
			Attrs: map[string]any{
				"label":     "Report Deceptive Site…",
				"accesskey": "D",
			},
		},
		{
			Key: "menu-help-not-deceptive",
			Attrs: map[string]any{
				"label":     "This Isn’t a Deceptive Site…",
				"accesskey": "D",
			},
		},
		{Key: "browser-window-minimize-button", Attrs: map[string]any{"tooltiptext": "Minimize"}},
		{Key: "browser-window-maximize-button", Attrs: map[string]any{"tooltiptext": "Maximize"}},
		{Key: "browser-window-restore-down-button", Attrs: map[string]any{"tooltiptext": "Restore Down"}},
		{Key: "browser-window-close-button", Attrs: map[string]any{"tooltiptext": "Close"}},
		{Key: "browser-window-minimize-button", Attrs: map[string]any{"tooltiptext": "Minimize"}},
		{Key: "browser-window-maximize-button", Attrs: map[string]any{"tooltiptext": "Maximize"}},
		{Key: "browser-window-restore-down-button", Attrs: map[string]any{"tooltiptext": "Restore Down"}},
		{Key: "browser-window-close-button", Attrs: map[string]any{"tooltiptext": "Close"}},
		{Key: "toolbar-button-back", Attrs: map[string]any{"label": "Back"}},
		{Key: "toolbar-button-forward", Attrs: map[string]any{"label": "Forward"}},
		{Key: "toolbar-button-stop-reload", Attrs: map[string]any{"title": "Reload"}},
		{Key: "toolbar-button-reload", Attrs: map[string]any{"label": "Reload"}},
		{Key: "toolbar-button-stop", Attrs: map[string]any{"label": "Stop"}},
		{Key: "urlbar-identity-button", Attrs: map[string]any{"aria-label": "View site information"}},
		{Key: "urlbar-permissions-granted", Attrs: map[string]any{"tooltiptext": "You have granted this website additional permissions."}},
		{Key: "urlbar-geolocation-blocked", Attrs: map[string]any{"tooltiptext": "You have blocked location information for this website."}},
		{Key: "urlbar-xr-blocked", Attrs: map[string]any{"tooltiptext": "You have blocked virtual reality device access for this website."}},
		{Key: "urlbar-web-notifications-blocked", Attrs: map[string]any{"tooltiptext": "You have blocked notifications for this website."}},
		{Key: "urlbar-camera-blocked", Attrs: map[string]any{"tooltiptext": "You have blocked your camera for this website."}},
		{Key: "urlbar-microphone-blocked", Attrs: map[string]any{"tooltiptext": "You have blocked your microphone for this website."}},
		{Key: "urlbar-screen-blocked", Attrs: map[string]any{"tooltiptext": "You have blocked this website from sharing your screen."}},
		{Key: "urlbar-persistent-storage-blocked", Attrs: map[string]any{"tooltiptext": "You have blocked persistent storage for this website."}},
		{Key: "urlbar-popup-blocked", Attrs: map[string]any{"tooltiptext": "You have blocked pop-ups for this website."}},
		{
			Key: "urlbar-autoplay-media-blocked",
			Attrs: map[string]any{
				"tooltiptext": "You have blocked autoplay media with sound for this website.",
			},
		},
		{Key: "urlbar-canvas-blocked", Attrs: map[string]any{"tooltiptext": "You have blocked canvas data extraction for this website."}},
		{Key: "urlbar-midi-blocked", Attrs: map[string]any{"tooltiptext": "You have blocked MIDI access for this website."}},
		{Key: "urlbar-install-blocked", Attrs: map[string]any{"tooltiptext": "You have blocked add-on installation for this website."}},
		{Key: "urlbar-default-notification-anchor", Attrs: map[string]any{"tooltiptext": "Open message panel"}},
		{Key: "urlbar-geolocation-notification-anchor", Attrs: map[string]any{"tooltiptext": "Open location request panel"}},
		{Key: "urlbar-xr-notification-anchor", Attrs: map[string]any{"tooltiptext": "Open virtual reality permission panel"}},
		{Key: "urlbar-autoplay-notification-anchor", Attrs: map[string]any{"tooltiptext": "Open autoplay panel"}},
		{Key: "urlbar-addons-notification-anchor", Attrs: map[string]any{"tooltiptext": "Open add-on installation message panel"}},
		{Key: "urlbar-canvas-notification-anchor", Attrs: map[string]any{"tooltiptext": "Manage canvas extraction permission"}},
		{Key: "urlbar-indexed-db-notification-anchor", Attrs: map[string]any{"tooltiptext": "Open offline storage message panel"}},
		{Key: "urlbar-password-notification-anchor", Attrs: map[string]any{"tooltiptext": "Open save password message panel"}},
		{Key: "urlbar-plugins-notification-anchor", Attrs: map[string]any{"tooltiptext": "Manage plug-in use"}},
		{Key: "urlbar-web-notification-anchor", Attrs: map[string]any{"tooltiptext": "Change whether you can receive notifications from the site"}},
		{
			Key: "urlbar-web-rtc-share-devices-notification-anchor",
			Attrs: map[string]any{
				"tooltiptext": "Manage sharing your camera and/or microphone with the site",
			},
		},
		{
			Key: "urlbar-web-rtc-share-microphone-notification-anchor",
			Attrs: map[string]any{
				"tooltiptext": "Manage sharing your microphone with the site",
			},
		},
		{
			Key: "urlbar-web-rtc-share-screen-notification-anchor",
			Attrs: map[string]any{
				"tooltiptext": "Manage sharing your windows or screen with the site",
			},
		},
		{Key: "urlbar-services-notification-anchor", Attrs: map[string]any{"tooltiptext": "Open install message panel"}},
		{Key: "urlbar-translate-notification-anchor", Attrs: map[string]any{"tooltiptext": "Translate this page"}},
		{Key: "urlbar-translated-notification-anchor", Attrs: map[string]any{"tooltiptext": "Manage page translation"}},
		{Key: "urlbar-eme-notification-anchor", Attrs: map[string]any{"tooltiptext": "Manage use of DRM software"}},
		{Key: "urlbar-persistent-storage-notification-anchor", Attrs: map[string]any{"tooltiptext": "Store data in Persistent Storage"}},
		{Key: "urlbar-midi-notification-anchor", Attrs: map[string]any{"tooltiptext": "Open MIDI panel"}},
		{Key: "urlbar-web-authn-anchor", Attrs: map[string]any{"tooltiptext": "Open Web Authentication panel"}},
		{Key: "urlbar-storage-access-anchor", Attrs: map[string]any{"tooltiptext": "Open browsing activity permission panel"}},
		{Key: "urlbar-remote-control-notification-anchor", Attrs: map[string]any{"tooltiptext": "Browser is under remote control"}},
		{Key: "urlbar-switch-to-tab", Attrs: map[string]any{"value": "Switch to tab:"}},
		{Key: "urlbar-extension", Attrs: map[string]any{"value": "Extension:"}},
		{Key: "urlbar-placeholder", Attrs: map[string]any{"placeholder": "Search or enter address"}},
		{Key: "urlbar-go-button", Attrs: map[string]any{"tooltiptext": "Go to the address in the Location Bar"}},
		{Key: "urlbar-page-action-button", Attrs: map[string]any{"tooltiptext": "Page actions"}},
		{Key: "urlbar-pocket-button", Attrs: map[string]any{"tooltiptext": "Save to Pocket"}},
		{Key: "browser-window-minimize-button", Attrs: map[string]any{"tooltiptext": "Minimize"}},
		{Key: "browser-window-restore-down-button", Attrs: map[string]any{"tooltiptext": "Restore Down"}},
		{Key: "browser-window-close-button", Attrs: map[string]any{"tooltiptext": "Close"}},
		{
			Key: "bookmarks-toolbar",
			Attrs: map[string]any{
				"toolbarname": "Bookmarks Toolbar",
				"accesskey":   "B",
				"aria-label":  "Bookmarks",
			},
		},
		{
			Key:   "bookmarks-toolbar-empty-message",
			Value: ptr.Ptr("For quick access, place your bookmarks here on the bookmarks toolbar. <a data-l10n-name=\"manage-bookmarks\">Manage bookmarks…</a>"),
		},
		{Key: "bookmarks-toolbar-placeholder", Attrs: map[string]any{"title": "Bookmarks Toolbar Items"}},
		{Key: "bookmarks-toolbar-placeholder-button", Attrs: map[string]any{"label": "Bookmarks Toolbar Items"}},
		{Key: "bookmarks-toolbar-placeholder-button", Attrs: map[string]any{"label": "Bookmarks Toolbar Items"}},
		{Key: "bookmarks-toolbar-chevron", Attrs: map[string]any{"tooltiptext": "Show more bookmarks"}},
		{Key: "fullscreen-warning-no-domain", Value: ptr.Ptr("This document is now full screen")},
		{Key: "fullscreen-exit-button", Value: ptr.Ptr("Exit Full Screen (Esc)")},
		{Key: "pointerlock-warning-no-domain", Value: ptr.Ptr("This document has control of your pointer. Press Esc to take back control.")},
		{
			Key: "text-action-undo",
			Attrs: map[string]any{
				"label":     "Undo",
				"accesskey": "U",
			},
		},
		{
			Key: "text-action-cut",
			Attrs: map[string]any{
				"label":     "Cut",
				"accesskey": "t",
			},
		},
		{
			Key: "text-action-copy",
			Attrs: map[string]any{
				"label":     "Copy",
				"accesskey": "C",
			},
		},
		{
			Key: "text-action-paste",
			Attrs: map[string]any{
				"label":     "Paste",
				"accesskey": "P",
			},
		},
		{
			Key: "text-action-delete",
			Attrs: map[string]any{
				"label":     "Delete",
				"accesskey": "D",
			},
		},
		{
			Key: "text-action-select-all",
			Attrs: map[string]any{
				"label":     "Select All",
				"accesskey": "A",
			},
		},
		{
			Key: "urlbar-pocket-button",
			Args: map[string]any{
				"tabCount": "1",
			},
			Attrs: map[string]any{
				"tooltiptext": "Save to Pocket",
			},
		},
		{Key: "menu-bookmark-this-page", Attrs: map[string]any{"label": "Bookmark This Page"}},
		{
			Key: "urlbar-star-add-bookmark",
			Args: map[string]any{
				"shortcut": "Ctrl+D",
			},
			Attrs: map[string]any{
				"tooltiptext": "Bookmark this page (Ctrl+D)",
			},
		},
		{
			Key: "main-context-menu-bookmark-add-with-shortcut",
			Args: map[string]any{
				"shortcut": "Ctrl+D",
			},
			Attrs: map[string]any{
				"aria-label":  "Bookmark This Page",
				"accesskey":   "m",
				"tooltiptext": "Bookmark this page (Ctrl+D)",
			},
		},
	},
}

var PreferencesScenario = Scenario{
	Name:   "preferences",
	Locale: "en-US",
	FileSources: []FileSource{
		{Name: "toolkit", PathScheme: "toolkit/{locale}/"},
		{Name: "browser", PathScheme: "browser/{locale}/"},
	},
	Queries: []Query{
		{Key: "pref-page-title", Value: ptr.Ptr("Preferences")},
		{Key: "category-list", Attrs: map[string]any{"aria-label": "Categories"}},
		{Key: "category-general", Attrs: map[string]any{"tooltiptext": "General"}},
		{Key: "pane-general-title", Value: ptr.Ptr("General")},
		{Key: "category-home", Attrs: map[string]any{"tooltiptext": "Home"}},
		{Key: "pane-home-title", Value: ptr.Ptr("Home")},
		{Key: "category-search", Attrs: map[string]any{"tooltiptext": "Search"}},
		{Key: "pane-search-title", Value: ptr.Ptr("Search")},
		{Key: "category-privacy", Attrs: map[string]any{"tooltiptext": "Privacy & Security"}},
		{Key: "pane-privacy-title", Value: ptr.Ptr("Privacy & Security")},
		{Key: "category-sync2", Attrs: map[string]any{"tooltiptext": "Sync"}},
		{Key: "pane-sync-title2", Value: ptr.Ptr("Sync")},
		{Key: "category-experimental", Attrs: map[string]any{"tooltiptext": "Nightly Experiments"}},
		{Key: "pane-experimental-title", Value: ptr.Ptr("Nightly Experiments")},
		{Key: "addons-button-label", Value: ptr.Ptr("Extensions & Themes")},
		{Key: "help-button-label", Value: ptr.Ptr("Nightly Support")},
		{Key: "focus-search", Attrs: map[string]any{"key": "f"}},
		{Key: "managed-notice", Value: ptr.Ptr("Your browser is being managed by your organization.")},
		{
			Key: "search-input-box",
			Attrs: map[string]any{
				"style":       "width: 15.4em",
				"placeholder": "Find in Preferences",
			},
		},
		{Key: "search-results-header", Value: ptr.Ptr("Search Results")},
		{
			Key:   "search-results-empty-message",
			Value: ptr.Ptr("Sorry! There are no results in Preferences for “<span data-l10n-name=\"query\"></span>”."),
		},
		{Key: "search-results-help-link", Value: ptr.Ptr("Need help? Visit <a data-l10n-name=\"url\">Nightly Support</a>")},
		{Key: "containers-back-button", Attrs: map[string]any{"aria-label": "Back to Preferences"}},
		{Key: "containers-header", Value: ptr.Ptr("Container Tabs")},
		{
			Key: "containers-add-button",
			Attrs: map[string]any{
				"label":     "Add New Container",
				"accesskey": "A",
			},
		},
		{
			Key: "containers-new-tab-check",
			Attrs: map[string]any{
				"label":     "Select a container for each new tab",
				"accesskey": "S",
			},
		},
		{Key: "close-button", Attrs: map[string]any{"aria-label": "Close"}},
		{Key: "pane-general-title", Value: ptr.Ptr("General")},
		{Key: "startup-header", Value: ptr.Ptr("Startup")},
		{
			Key: "startup-restore-previous-session",
			Attrs: map[string]any{
				"label":     "Restore previous session",
				"accesskey": "s",
			},
		},
		{Key: "startup-restore-warn-on-quit", Attrs: map[string]any{"label": "Warn you when quitting the browser"}},
		{
			Key: "always-check-default",
			Attrs: map[string]any{
				"label":     "Always check if Nightly is your default browser",
				"accesskey": "y",
			},
		},
		{Key: "is-not-default", Value: ptr.Ptr("Nightly is not your default browser")},
		{
			Key: "set-as-my-default-browser",
			Attrs: map[string]any{
				"label":     "Make Default…",
				"accesskey": "D",
			},
		},
		{Key: "is-default", Value: ptr.Ptr("Nightly is currently your default browser")},
		{Key: "tabs-group-header", Value: ptr.Ptr("Tabs")},
		{
			Key: "ctrl-tab-recently-used-order",
			Attrs: map[string]any{
				"label":     "Ctrl+Tab cycles through tabs in recently used order",
				"accesskey": "T",
			},
		},
		{
			Key: "open-new-link-as-tabs",
			Attrs: map[string]any{
				"label":     "Open links in tabs instead of new windows",
				"accesskey": "w",
			},
		},
		{
			Key: "warn-on-close-multiple-tabs",
			Attrs: map[string]any{
				"label":     "Warn you when closing multiple tabs",
				"accesskey": "m",
			},
		},
		{
			Key: "warn-on-open-many-tabs",
			Attrs: map[string]any{
				"label":     "Warn you when opening multiple tabs might slow down Nightly",
				"accesskey": "d",
			},
		},
		{
			Key: "switch-links-to-new-tabs",
			Attrs: map[string]any{
				"label":     "When you open a link in a new tab, switch to it immediately",
				"accesskey": "h",
			},
		},
		{Key: "disable-extension", Attrs: map[string]any{"label": "Disable Extension"}},
		{
			Key: "browser-containers-enabled",
			Attrs: map[string]any{
				"label":     "Enable Container Tabs",
				"accesskey": "n",
			},
		},
		{Key: "browser-containers-learn-more", Value: ptr.Ptr("Learn more")},
		{
			Key: "browser-containers-settings",
			Attrs: map[string]any{
				"label":     "Settings…",
				"accesskey": "i",
			},
		},
		{Key: "language-and-appearance-header", Value: ptr.Ptr("Language and Appearance")},
		{Key: "fonts-and-colors-header", Value: ptr.Ptr("Fonts and Colors")},
		{
			Key:   "default-font",
			Value: ptr.Ptr("Default font"),
			Attrs: map[string]any{
				"accesskey": "D",
			},
		},
		{
			Key:   "default-font-size",
			Value: ptr.Ptr("Size"),
			Attrs: map[string]any{
				"accesskey": "S",
			},
		},
		{
			Key: "advanced-fonts",
			Attrs: map[string]any{
				"label":     "Advanced…",
				"accesskey": "A",
			},
		},
		{
			Key: "colors-settings",
			Attrs: map[string]any{
				"label":     "Colors…",
				"accesskey": "C",
			},
		},
		{Key: "preferences-zoom-header", Value: ptr.Ptr("Zoom")},
		{
			Key:   "preferences-default-zoom",
			Value: ptr.Ptr("Default zoom"),
			Attrs: map[string]any{
				"accesskey": "z",
			},
		},
		{
			Key: "preferences-zoom-text-only",
			Attrs: map[string]any{
				"label":     "Zoom text only",
				"accesskey": "t",
			},
		},
		{Key: "language-header", Value: ptr.Ptr("Language")},
		{
			Key:   "choose-browser-language-description",
			Value: ptr.Ptr("Choose the languages used to display menus, messages, and notifications from Nightly."),
		},
		{
			Key: "manage-browser-languages-button",
			Attrs: map[string]any{
				"label":     "Set Alternatives…",
				"accesskey": "l",
			},
		},
		{Key: "choose-language-description", Value: ptr.Ptr("Choose your preferred language for displaying pages")},
		{
			Key: "choose-button",
			Attrs: map[string]any{
				"label":     "Choose…",
				"accesskey": "o",
			},
		},
		{
			Key: "use-system-locale",
			Args: map[string]any{
				"localeName": "und",
			},
			Attrs: map[string]any{
				"label": "Use your operating system settings for “und” to format dates, times, numbers, and measurements.",
			},
		},
		{
			Key: "translate-web-pages",
			Attrs: map[string]any{
				"label":     "Translate web content",
				"accesskey": "T",
			},
		},
		{Key: "translate-attribution", Value: ptr.Ptr("Translations by <img data-l10n-name=\"logo\"/>")},
		{
			Key: "translate-exceptions",
			Attrs: map[string]any{
				"label":     "Exceptions…",
				"accesskey": "x",
			},
		},
		{
			Key: "check-user-spelling",
			Attrs: map[string]any{
				"label":     "Check your spelling as you type",
				"accesskey": "t",
			},
		},
		{Key: "files-and-applications-title", Value: ptr.Ptr("Files and Applications")},
		{Key: "download-header", Value: ptr.Ptr("Downloads")},
		{
			Key: "download-save-to",
			Attrs: map[string]any{
				"label":     "Save files to",
				"accesskey": "v",
			},
		},
		{
			Key: "download-choose-folder",
			Attrs: map[string]any{
				"label":     "Browse…",
				"accesskey": "o",
			},
		},
		{
			Key: "download-always-ask-where",
			Attrs: map[string]any{
				"label":     "Always ask you where to save files",
				"accesskey": "A",
			},
		},
		{Key: "applications-header", Value: ptr.Ptr("Applications")},
		{
			Key:   "applications-description",
			Value: ptr.Ptr("Choose how Nightly handles the files you download from the web or the applications you use while browsing."),
		},
		{Key: "applications-filter", Attrs: map[string]any{"placeholder": "Search file types or applications"}},
		{
			Key: "applications-type-column",
			Attrs: map[string]any{
				"label":     "Content Type",
				"accesskey": "T",
			},
		},
		{
			Key: "applications-action-column",
			Attrs: map[string]any{
				"label":     "Action",
				"accesskey": "A",
			},
		},
		{Key: "drm-content-header", Value: ptr.Ptr("Digital Rights Management (DRM) Content")},
		{
			Key: "play-drm-content",
			Attrs: map[string]any{
				"label":     "Play DRM-controlled content",
				"accesskey": "P",
			},
		},
		{Key: "play-drm-content-learn-more", Value: ptr.Ptr("Learn more")},
		{Key: "update-application-title", Value: ptr.Ptr("Nightly Updates")},
		{Key: "update-application-title", Value: ptr.Ptr("Nightly Updates")},
		{Key: "update-application-description", Value: ptr.Ptr("Keep Nightly up to date for the best performance, stability, and security.")},
		{
			Key: "update-history",
			Attrs: map[string]any{
				"label":     "Show Update History…",
				"accesskey": "p",
			},
		},
		{
			Key: "update-checkForUpdatesButton",
			Attrs: map[string]any{
				"label":     "Check for updates",
				"accesskey": "C",
			},
		},
		{
			Key: "update-updateButton",
			Attrs: map[string]any{
				"label":     "Restart to Update Nightly",
				"accesskey": "R",
			},
		},
		{Key: "update-checkingForUpdates", Value: ptr.Ptr("Checking for updates…")},
		{
			Key: "update-checkForUpdatesButton",
			Attrs: map[string]any{
				"label":     "Check for updates",
				"accesskey": "C",
			},
		},
		{
			Key:   "update-downloading",
			Value: ptr.Ptr("<img data-l10n-name=\"icon\"/>Downloading update — <label data-l10n-name=\"download-status\"/>"),
		},
		{Key: "update-applying", Value: ptr.Ptr("Applying update…")},
		{
			Key:   "update-failed-main",
			Value: ptr.Ptr("Update failed. <a data-l10n-name=\"failed-link-main\">Download the latest version</a>"),
		},
		{
			Key: "update-checkForUpdatesButton",
			Attrs: map[string]any{
				"label":     "Check for updates",
				"accesskey": "C",
			},
		},
		{Key: "update-adminDisabled", Value: ptr.Ptr("Updates disabled by your system administrator")},
		{
			Key: "update-checkForUpdatesButton",
			Attrs: map[string]any{
				"label":     "Check for updates",
				"accesskey": "C",
			},
		},
		{Key: "update-noUpdatesFound", Value: ptr.Ptr("Nightly is up to date")},
		{
			Key: "update-checkForUpdatesButton",
			Attrs: map[string]any{
				"label":     "Check for updates",
				"accesskey": "C",
			},
		},
		{Key: "update-otherInstanceHandlingUpdates", Value: ptr.Ptr("Nightly is being updated by another instance")},
		{
			Key: "update-checkForUpdatesButton",
			Attrs: map[string]any{
				"label":     "Check for updates",
				"accesskey": "C",
			},
		},
		{Key: "update-manual", Value: ptr.Ptr("Updates available at <label data-l10n-name=\"manual-link\"/>")},
		{
			Key: "update-checkForUpdatesButton",
			Attrs: map[string]any{
				"label":     "Check for updates",
				"accesskey": "C",
			},
		},
		{
			Key:   "update-unsupported",
			Value: ptr.Ptr("You can not perform further updates on this system. <label data-l10n-name=\"unsupported-link\">Learn more</label>"),
		},
		{
			Key: "update-checkForUpdatesButton",
			Attrs: map[string]any{
				"label":     "Check for updates",
				"accesskey": "C",
			},
		},
		{Key: "update-restarting", Value: ptr.Ptr("Restarting…")},
		{
			Key: "update-updateButton",
			Attrs: map[string]any{
				"label":     "Restart to Update Nightly",
				"accesskey": "R",
			},
		},
		{Key: "update-application-allow-description", Value: ptr.Ptr("Allow Nightly to")},
		{
			Key: "update-application-auto",
			Attrs: map[string]any{
				"label":     "Automatically install updates (recommended)",
				"accesskey": "A",
			},
		},
		{
			Key: "update-application-check-choose",
			Attrs: map[string]any{
				"label":     "Check for updates but let you choose to install them",
				"accesskey": "C",
			},
		},
		{
			Key:   "update-application-warning-cross-user-setting",
			Value: ptr.Ptr("This setting will apply to all Windows accounts and Nightly profiles using this installation of Nightly."),
		},
		{Key: "performance-title", Value: ptr.Ptr("Performance")},
		{Key: "performance-title", Value: ptr.Ptr("Performance")},
		{
			Key: "performance-use-recommended-settings-checkbox",
			Attrs: map[string]any{
				"label":     "Use recommended performance settings",
				"accesskey": "U",
			},
		},
		{Key: "performance-settings-learn-more", Value: ptr.Ptr("Learn more")},
		{Key: "performance-use-recommended-settings-desc", Value: ptr.Ptr("These settings are tailored to your computer’s hardware and operating system.")},
		{
			Key: "performance-allow-hw-accel",
			Attrs: map[string]any{
				"label":     "Use hardware acceleration when available",
				"accesskey": "r",
			},
		},
		{
			Key:   "performance-limit-content-process-option",
			Value: ptr.Ptr("Content process limit"),
			Attrs: map[string]any{
				"accesskey": "l",
			},
		},
		{
			Key:   "performance-limit-content-process-enabled-desc",
			Value: ptr.Ptr("Additional content processes can improve performance when using multiple tabs, but will also use more memory."),
		},
		{
			Key:   "performance-limit-content-process-blocked-desc",
			Value: ptr.Ptr("Modifying the number of content processes is only possible with multiprocess Nightly. <a data-l10n-name=\"learn-more\">Learn how to check if multiprocess is enabled</a>"),
		},
		{Key: "browsing-title", Value: ptr.Ptr("Browsing")},
		{Key: "browsing-title", Value: ptr.Ptr("Browsing")},
		{
			Key: "browsing-use-autoscroll",
			Attrs: map[string]any{
				"label":     "Use autoscrolling",
				"accesskey": "a",
			},
		},
		{
			Key: "browsing-use-smooth-scrolling",
			Attrs: map[string]any{
				"label":     "Use smooth scrolling",
				"accesskey": "m",
			},
		},
		{
			Key: "browsing-use-cursor-navigation",
			Attrs: map[string]any{
				"label":     "Always use the cursor keys to navigate within pages",
				"accesskey": "k",
			},
		},
		{
			Key: "browsing-search-on-start-typing",
			Attrs: map[string]any{
				"label":     "Search for text when you start typing",
				"accesskey": "x",
			},
		},
		{
			Key: "browsing-picture-in-picture-toggle-enabled",
			Attrs: map[string]any{
				"label":     "Enable picture-in-picture video controls",
				"accesskey": "E",
			},
		},
		{Key: "browsing-picture-in-picture-learn-more", Value: ptr.Ptr("Learn more")},
		{
			Key: "browsing-media-control",
			Attrs: map[string]any{
				"label":     "Control media via keyboard, headset, or virtual interface",
				"accesskey": "v",
			},
		},
		{Key: "browsing-media-control-learn-more", Value: ptr.Ptr("Learn more")},
		{
			Key: "browsing-cfr-recommendations",
			Attrs: map[string]any{
				"label":     "Recommend extensions as you browse",
				"accesskey": "R",
			},
		},
		{Key: "browsing-cfr-recommendations-learn-more", Value: ptr.Ptr("Learn more")},
		{
			Key: "browsing-cfr-features",
			Attrs: map[string]any{
				"label":     "Recommend features as you browse",
				"accesskey": "f",
			},
		},
		{Key: "browsing-cfr-recommendations-learn-more", Value: ptr.Ptr("Learn more")},
		{Key: "network-settings-title", Value: ptr.Ptr("Network Settings")},
		{Key: "network-settings-title", Value: ptr.Ptr("Network Settings")},
		{Key: "network-proxy-connection-learn-more", Value: ptr.Ptr("Learn more")},
		{
			Key: "network-proxy-connection-settings",
			Attrs: map[string]any{
				"label":     "Settings…",
				"accesskey": "e",
			},
		},
		{
			Key: "performance-default-content-process-count",
			Args: map[string]any{
				"num": "8",
			},
			Attrs: map[string]any{
				"label": "8 (default)",
			},
		},
		{
			Key: "update-application-version",
			Args: map[string]any{
				"version": "86.0a1 (2020-12-27) (64-bit)",
			},
			Value: ptr.Ptr("Version 86.0a1 (2020-12-27) (64-bit) <a data-l10n-name=\"learn-more\">What’s new</a>"),
		},
		{
			Key: "preferences-default-zoom-value",
			Args: map[string]any{
				"percentage": "30",
			},
			Attrs: map[string]any{
				"label": "30%",
			},
		},
		{
			Key: "preferences-default-zoom-value",
			Args: map[string]any{
				"percentage": "50",
			},
			Attrs: map[string]any{
				"label": "50%",
			},
		},
		{
			Key: "preferences-default-zoom-value",
			Args: map[string]any{
				"percentage": "67",
			},
			Attrs: map[string]any{
				"label": "67%",
			},
		},
		{
			Key: "preferences-default-zoom-value",
			Args: map[string]any{
				"percentage": "80",
			},
			Attrs: map[string]any{
				"label": "80%",
			},
		},
		{
			Key: "preferences-default-zoom-value",
			Args: map[string]any{
				"percentage": "90",
			},
			Attrs: map[string]any{
				"label": "90%",
			},
		},
		{
			Key: "preferences-default-zoom-value",
			Args: map[string]any{
				"percentage": "100",
			},
			Attrs: map[string]any{
				"label": "100%",
			},
		},
		{
			Key: "preferences-default-zoom-value",
			Args: map[string]any{
				"percentage": "110",
			},
			Attrs: map[string]any{
				"label": "110%",
			},
		},
		{
			Key: "preferences-default-zoom-value",
			Args: map[string]any{
				"percentage": "120",
			},
			Attrs: map[string]any{
				"label": "120%",
			},
		},
		{
			Key: "preferences-default-zoom-value",
			Args: map[string]any{
				"percentage": "133",
			},
			Attrs: map[string]any{
				"label": "133%",
			},
		},
		{
			Key: "preferences-default-zoom-value",
			Args: map[string]any{
				"percentage": "150",
			},
			Attrs: map[string]any{
				"label": "150%",
			},
		},
		{
			Key: "preferences-default-zoom-value",
			Args: map[string]any{
				"percentage": "170",
			},
			Attrs: map[string]any{
				"label": "170%",
			},
		},
		{
			Key: "preferences-default-zoom-value",
			Args: map[string]any{
				"percentage": "200",
			},
			Attrs: map[string]any{
				"label": "200%",
			},
		},
		{
			Key: "preferences-default-zoom-value",
			Args: map[string]any{
				"percentage": "240",
			},
			Attrs: map[string]any{
				"label": "240%",
			},
		},
		{
			Key: "preferences-default-zoom-value",
			Args: map[string]any{
				"percentage": "300",
			},
			Attrs: map[string]any{
				"label": "300%",
			},
		},
		{
			Key: "preferences-default-zoom-value",
			Args: map[string]any{
				"percentage": "400",
			},
			Attrs: map[string]any{
				"label": "400%",
			},
		},
		{
			Key: "preferences-default-zoom-value",
			Args: map[string]any{
				"percentage": "500",
			},
			Attrs: map[string]any{
				"label": "500%",
			},
		},
		{Key: "network-proxy-connection-description", Value: ptr.Ptr("Configure how Nightly connects to the internet.")},
		{
			Key: "fonts-label-default",
			Args: map[string]any{
				"name": "DejaVu Serif",
			},
			Attrs: map[string]any{
				"label": "Default (DejaVu Serif)",
			},
		},
		{Key: "applications-open-inapp-label", Attrs: map[string]any{"value": "Open in Nightly"}},
		{Key: "applications-open-inapp", Attrs: map[string]any{"label": "Open in Nightly"}},
		{Key: "applications-open-inapp", Attrs: map[string]any{"label": "Open in Nightly"}},
		{Key: "applications-always-ask", Attrs: map[string]any{"label": "Always ask"}},
		{Key: "applications-action-save", Attrs: map[string]any{"label": "Save File"}},
		{Key: "applications-use-os-default", Attrs: map[string]any{"label": "Use system default application"}},
		{Key: "applications-use-other", Attrs: map[string]any{"label": "Use other…"}},
		{Key: "applications-open-inapp-label", Attrs: map[string]any{"value": "Open in Nightly"}},
		{Key: "applications-open-inapp", Attrs: map[string]any{"label": "Open in Nightly"}},
		{Key: "applications-open-inapp", Attrs: map[string]any{"label": "Open in Nightly"}},
		{Key: "applications-always-ask", Attrs: map[string]any{"label": "Always ask"}},
		{Key: "applications-action-save", Attrs: map[string]any{"label": "Save File"}},
		{Key: "applications-use-os-default", Attrs: map[string]any{"label": "Use system default application"}},
		{Key: "applications-use-other", Attrs: map[string]any{"label": "Use other…"}},
		{Key: "applications-open-inapp-label", Attrs: map[string]any{"value": "Open in Nightly"}},
		{Key: "applications-open-inapp", Attrs: map[string]any{"label": "Open in Nightly"}},
		{Key: "applications-open-inapp", Attrs: map[string]any{"label": "Open in Nightly"}},
		{Key: "applications-always-ask", Attrs: map[string]any{"label": "Always ask"}},
		{Key: "applications-action-save", Attrs: map[string]any{"label": "Save File"}},
		{Key: "applications-use-os-default", Attrs: map[string]any{"label": "Use system default application"}},
		{Key: "applications-use-other", Attrs: map[string]any{"label": "Use other…"}},
		{Key: "applications-open-inapp-label", Attrs: map[string]any{"value": "Open in Nightly"}},
		{Key: "applications-open-inapp", Attrs: map[string]any{"label": "Open in Nightly"}},
		{Key: "applications-open-inapp", Attrs: map[string]any{"label": "Open in Nightly"}},
		{Key: "applications-always-ask", Attrs: map[string]any{"label": "Always ask"}},
		{Key: "applications-action-save", Attrs: map[string]any{"label": "Save File"}},
		{Key: "applications-use-os-default", Attrs: map[string]any{"label": "Use system default application"}},
		{Key: "applications-use-other", Attrs: map[string]any{"label": "Use other…"}},
		{Key: "applications-always-ask-label", Attrs: map[string]any{"value": "Always ask"}},
		{Key: "applications-always-ask", Attrs: map[string]any{"label": "Always ask"}},
		{Key: "applications-always-ask", Attrs: map[string]any{"label": "Always ask"}},
		{
			Key: "applications-use-app-default",
			Args: map[string]any{
				"app-name": "Polari",
			},
			Attrs: map[string]any{
				"label": "Use Polari (default)",
			},
		},
		{
			Key: "applications-use-app",
			Args: map[string]any{
				"app-name": "Mibbit",
			},
			Attrs: map[string]any{
				"label": "Use Mibbit",
			},
		},
		{Key: "applications-use-other", Attrs: map[string]any{"label": "Use other…"}},
		{Key: "applications-manage-app", Attrs: map[string]any{"label": "Application Details…"}},
		{Key: "applications-always-ask-label", Attrs: map[string]any{"value": "Always ask"}},
		{Key: "applications-always-ask", Attrs: map[string]any{"label": "Always ask"}},
		{Key: "applications-always-ask", Attrs: map[string]any{"label": "Always ask"}},
		{
			Key: "applications-use-app",
			Args: map[string]any{
				"app-name": "Mibbit",
			},
			Attrs: map[string]any{
				"label": "Use Mibbit",
			},
		},
		{Key: "applications-use-other", Attrs: map[string]any{"label": "Use other…"}},
		{Key: "applications-manage-app", Attrs: map[string]any{"label": "Application Details…"}},
		{
			Key: "applications-use-app-default-label",
			Args: map[string]any{
				"app-name": "Evolution",
			},
			Attrs: map[string]any{
				"value": "Use Evolution (default)",
			},
		},
		{
			Key: "applications-use-app-default",
			Args: map[string]any{
				"app-name": "Evolution",
			},
			Attrs: map[string]any{
				"label": "Use Evolution (default)",
			},
		},
		{Key: "applications-always-ask", Attrs: map[string]any{"label": "Always ask"}},
		{
			Key: "applications-use-app-default",
			Args: map[string]any{
				"app-name": "Evolution",
			},
			Attrs: map[string]any{
				"label": "Use Evolution (default)",
			},
		},
		{
			Key: "applications-use-app",
			Args: map[string]any{
				"app-name": "Yahoo! Mail",
			},
			Attrs: map[string]any{
				"label": "Use Yahoo! Mail",
			},
		},
		{
			Key: "applications-use-app",
			Args: map[string]any{
				"app-name": "Gmail",
			},
			Attrs: map[string]any{
				"label": "Use Gmail",
			},
		},
		{Key: "applications-use-other", Attrs: map[string]any{"label": "Use other…"}},
		{Key: "applications-manage-app", Attrs: map[string]any{"label": "Application Details…"}},
		{Key: "applications-open-inapp-label", Attrs: map[string]any{"value": "Open in Nightly"}},
		{Key: "applications-open-inapp", Attrs: map[string]any{"label": "Open in Nightly"}},
		{Key: "applications-open-inapp", Attrs: map[string]any{"label": "Open in Nightly"}},
		{Key: "applications-always-ask", Attrs: map[string]any{"label": "Always ask"}},
		{Key: "applications-action-save", Attrs: map[string]any{"label": "Save File"}},
		{Key: "applications-use-os-default", Attrs: map[string]any{"label": "Use system default application"}},
		{Key: "applications-use-other", Attrs: map[string]any{"label": "Use other…"}},
		{Key: "applications-open-inapp-label", Attrs: map[string]any{"value": "Open in Nightly"}},
		{Key: "applications-open-inapp", Attrs: map[string]any{"label": "Open in Nightly"}},
		{Key: "applications-open-inapp", Attrs: map[string]any{"label": "Open in Nightly"}},
		{Key: "applications-always-ask", Attrs: map[string]any{"label": "Always ask"}},
		{Key: "applications-action-save", Attrs: map[string]any{"label": "Save File"}},
		{Key: "applications-use-os-default", Attrs: map[string]any{"label": "Use system default application"}},
		{Key: "applications-use-other", Attrs: map[string]any{"label": "Use other…"}},
		{Key: "applications-open-inapp-label", Attrs: map[string]any{"value": "Open in Nightly"}},
		{Key: "applications-open-inapp", Attrs: map[string]any{"label": "Open in Nightly"}},
		{Key: "applications-open-inapp", Attrs: map[string]any{"label": "Open in Nightly"}},
		{Key: "applications-always-ask", Attrs: map[string]any{"label": "Always ask"}},
		{Key: "applications-action-save", Attrs: map[string]any{"label": "Save File"}},
		{Key: "applications-use-os-default", Attrs: map[string]any{"label": "Use system default application"}},
		{Key: "applications-use-other", Attrs: map[string]any{"label": "Use other…"}},
		{Key: "applications-open-inapp-label", Attrs: map[string]any{"value": "Open in Nightly"}},
		{Key: "applications-open-inapp", Attrs: map[string]any{"label": "Open in Nightly"}},
		{Key: "applications-open-inapp", Attrs: map[string]any{"label": "Open in Nightly"}},
		{Key: "applications-always-ask", Attrs: map[string]any{"label": "Always ask"}},
		{Key: "applications-action-save", Attrs: map[string]any{"label": "Save File"}},
		{Key: "applications-use-os-default", Attrs: map[string]any{"label": "Use system default application"}},
		{Key: "applications-use-other", Attrs: map[string]any{"label": "Use other…"}},
		{Key: "applications-always-ask-label", Attrs: map[string]any{"value": "Always ask"}},
		{Key: "applications-always-ask", Attrs: map[string]any{"label": "Always ask"}},
		{Key: "applications-always-ask", Attrs: map[string]any{"label": "Always ask"}},
		{
			Key: "applications-use-app-default",
			Args: map[string]any{
				"app-name": "Polari",
			},
			Attrs: map[string]any{
				"label": "Use Polari (default)",
			},
		},
		{
			Key: "applications-use-app",
			Args: map[string]any{
				"app-name": "Mibbit",
			},
			Attrs: map[string]any{
				"label": "Use Mibbit",
			},
		},
		{Key: "applications-use-other", Attrs: map[string]any{"label": "Use other…"}},
		{Key: "applications-manage-app", Attrs: map[string]any{"label": "Application Details…"}},
		{Key: "applications-always-ask-label", Attrs: map[string]any{"value": "Always ask"}},
		{Key: "applications-always-ask", Attrs: map[string]any{"label": "Always ask"}},
		{Key: "applications-always-ask", Attrs: map[string]any{"label": "Always ask"}},
		{
			Key: "applications-use-app",
			Args: map[string]any{
				"app-name": "Mibbit",
			},
			Attrs: map[string]any{
				"label": "Use Mibbit",
			},
		},
		{Key: "applications-use-other", Attrs: map[string]any{"label": "Use other…"}},
		{Key: "applications-manage-app", Attrs: map[string]any{"label": "Application Details…"}},
		{
			Key: "applications-use-app-default-label",
			Args: map[string]any{
				"app-name": "Evolution",
			},
			Attrs: map[string]any{
				"value": "Use Evolution (default)",
			},
		},
		{
			Key: "applications-use-app-default",
			Args: map[string]any{
				"app-name": "Evolution",
			},
			Attrs: map[string]any{
				"label": "Use Evolution (default)",
			},
		},
		{Key: "applications-always-ask", Attrs: map[string]any{"label": "Always ask"}},
		{
			Key: "applications-use-app-default",
			Args: map[string]any{
				"app-name": "Evolution",
			},
			Attrs: map[string]any{
				"label": "Use Evolution (default)",
			},
		},
		{
			Key: "applications-use-app",
			Args: map[string]any{
				"app-name": "Yahoo! Mail",
			},
			Attrs: map[string]any{
				"label": "Use Yahoo! Mail",
			},
		},
		{
			Key: "applications-use-app",
			Args: map[string]any{
				"app-name": "Gmail",
			},
			Attrs: map[string]any{
				"label": "Use Gmail",
			},
		},
		{Key: "applications-use-other", Attrs: map[string]any{"label": "Use other…"}},
		{Key: "applications-manage-app", Attrs: map[string]any{"label": "Application Details…"}},
		{Key: "pane-home-title", Value: ptr.Ptr("Home")},
		{
			Key: "home-restore-defaults",
			Attrs: map[string]any{
				"label":     "Restore Defaults",
				"accesskey": "R",
			},
		},
		{Key: "home-new-windows-tabs-header", Value: ptr.Ptr("New Windows and Tabs")},
		{Key: "home-new-windows-tabs-description2", Value: ptr.Ptr("Choose what you see when you open your homepage, new windows, and new tabs.")},
		{Key: "home-homepage-mode-label", Value: ptr.Ptr("Homepage and new windows")},
		{Key: "home-mode-choice-default", Attrs: map[string]any{"label": "Firefox Home (Default)"}},
		{Key: "home-mode-choice-custom", Attrs: map[string]any{"label": "Custom URLs…"}},
		{Key: "home-mode-choice-blank", Attrs: map[string]any{"label": "Blank Page"}},
		{Key: "home-homepage-custom-url", Attrs: map[string]any{"placeholder": "Paste a URL…"}},
		{
			Key: "use-current-pages",
			Args: map[string]any{
				"tabCount": "0",
			},
			Attrs: map[string]any{
				"label":     "Use Current Pages",
				"accesskey": "C",
			},
		},
		{
			Key: "choose-bookmark",
			Attrs: map[string]any{
				"label":     "Use Bookmark…",
				"accesskey": "B",
			},
		},
		{Key: "home-newtabs-mode-label", Value: ptr.Ptr("New tabs")},
		{Key: "home-mode-choice-default", Attrs: map[string]any{"label": "Firefox Home (Default)"}},
		{Key: "home-mode-choice-blank", Attrs: map[string]any{"label": "Blank Page"}},
		{Key: "pane-search-title", Value: ptr.Ptr("Search")},
		{Key: "search-bar-header", Value: ptr.Ptr("Search Bar")},
		{Key: "search-bar-hidden", Attrs: map[string]any{"label": "Use the address bar for search and navigation"}},
		{Key: "search-bar-shown", Attrs: map[string]any{"label": "Add search bar in toolbar"}},
		{Key: "search-engine-default-header", Value: ptr.Ptr("Default Search Engine")},
		{
			Key:   "search-engine-default-desc-2",
			Value: ptr.Ptr("This is your default search engine in the address bar and search bar. You can switch it at any time."),
		},
		{
			Key: "search-separate-default-engine",
			Attrs: map[string]any{
				"label":     "Use this search engine in Private Windows",
				"accesskey": "U",
			},
		},
		{Key: "search-engine-default-private-desc-2", Value: ptr.Ptr("Choose a different default search engine for Private Windows only")},
		{Key: "search-suggestions-header", Value: ptr.Ptr("Search Suggestions")},
		{Key: "search-suggestions-desc", Value: ptr.Ptr("Choose how suggestions from search engines appear.")},
		{
			Key: "search-suggestions-option",
			Attrs: map[string]any{
				"label":     "Provide search suggestions",
				"accesskey": "s",
			},
		},
		{
			Key: "search-show-suggestions-url-bar-option",
			Attrs: map[string]any{
				"label":     "Show search suggestions in address bar results",
				"accesskey": "l",
			},
		},
		{
			Key: "search-show-suggestions-above-history-option",
			Attrs: map[string]any{
				"label": "Show search suggestions ahead of browsing history in address bar results",
			},
		},
		{Key: "search-show-suggestions-private-windows", Attrs: map[string]any{"label": "Show search suggestions in Private Windows"}},
		{
			Key:   "search-suggestions-cant-show",
			Value: ptr.Ptr("Search suggestions will not be shown in location bar results because you have configured Nightly to never remember history."),
		},
		{Key: "suggestions-addressbar-settings-generic", Value: ptr.Ptr("Change preferences for other address bar suggestions")},
		{Key: "search-one-click-header2", Value: ptr.Ptr("Search Shortcuts")},
		{
			Key:   "search-one-click-desc",
			Value: ptr.Ptr("Choose the alternative search engines that appear below the address bar and search bar when you start to enter a keyword."),
		},
		{Key: "search-choose-engine-column", Attrs: map[string]any{"label": "Search Engine"}},
		{Key: "search-choose-keyword-column", Attrs: map[string]any{"label": "Keyword"}},
		{
			Key: "search-restore-default",
			Attrs: map[string]any{
				"label":     "Restore Default Search Engines",
				"accesskey": "D",
			},
		},
		{
			Key: "search-remove-engine",
			Attrs: map[string]any{
				"label":     "Remove",
				"accesskey": "R",
			},
		},
		{
			Key: "search-add-engine",
			Attrs: map[string]any{
				"label":     "Add",
				"accesskey": "A",
			},
		},
		{Key: "search-find-more-link", Value: ptr.Ptr("Find more search engines")},
		{Key: "privacy-header", Value: ptr.Ptr("Browser Privacy")},
		{Key: "content-blocking-enhanced-tracking-protection", Value: ptr.Ptr("Enhanced Tracking Protection")},
		{
			Key:   "content-blocking-section-top-level-description",
			Value: ptr.Ptr("Trackers follow you around online to collect information about your browsing habits and interests. Nightly blocks many of these trackers and other malicious scripts."),
		},
		{Key: "content-blocking-learn-more", Value: ptr.Ptr("Learn more")},
		{
			Key: "tracking-manage-exceptions",
			Attrs: map[string]any{
				"label":     "Manage Exceptions…",
				"accesskey": "x",
			},
		},
		{
			Key:   "content-blocking-fpi-incompatibility-warning",
			Value: ptr.Ptr("You are using First Party Isolation (FPI), which overrides some of Nightly’s cookie settings."),
		},
		{
			Key: "enhanced-tracking-protection-setting-standard",
			Attrs: map[string]any{
				"label":     "Standard",
				"accesskey": "d",
			},
		},
		{Key: "content-blocking-expand-section", Attrs: map[string]any{"tooltiptext": "More information"}},
		{Key: "content-blocking-etp-standard-desc", Value: ptr.Ptr("Balanced for protection and performance. Pages will load normally.")},
		{Key: "content-blocking-social-media-trackers", Value: ptr.Ptr("Social media trackers")},
		{Key: "content-blocking-cross-site-cookies", Value: ptr.Ptr("Cross-site cookies")},
		{Key: "content-blocking-cross-site-tracking-cookies", Value: ptr.Ptr("Cross-site tracking cookies")},
		{Key: "content-blocking-cross-site-tracking-cookies-plus-isolate", Value: ptr.Ptr("Cross-site tracking cookies, and isolate remaining cookies")},
		{Key: "content-blocking-private-windows", Value: ptr.Ptr("Tracking content in Private Windows")},
		{Key: "content-blocking-all-windows-tracking-content", Value: ptr.Ptr("Tracking content in all windows")},
		{Key: "content-blocking-all-third-party-cookies", Value: ptr.Ptr("All third-party cookies")},
		{Key: "content-blocking-all-cookies", Value: ptr.Ptr("All cookies")},
		{Key: "content-blocking-unvisited-cookies", Value: ptr.Ptr("Cookies from unvisited sites")},
		{Key: "content-blocking-cryptominers", Value: ptr.Ptr("Cryptominers")},
		{Key: "content-blocking-fingerprinters", Value: ptr.Ptr("Fingerprinters")},
		{Key: "content-blocking-reload-description", Value: ptr.Ptr("You will need to reload your tabs to apply these changes.")},
		{
			Key: "content-blocking-reload-tabs-button",
			Attrs: map[string]any{
				"label":     "Reload All Tabs",
				"accesskey": "R",
			},
		},
		{
			Key: "enhanced-tracking-protection-setting-strict",
			Attrs: map[string]any{
				"label":     "Strict",
				"accesskey": "r",
			},
		},
		{Key: "content-blocking-expand-section", Attrs: map[string]any{"tooltiptext": "More information"}},
		{Key: "content-blocking-etp-strict-desc", Value: ptr.Ptr("Stronger protection, but may cause some sites or content to break.")},
		{Key: "content-blocking-social-media-trackers", Value: ptr.Ptr("Social media trackers")},
		{Key: "content-blocking-cross-site-tracking-cookies", Value: ptr.Ptr("Cross-site tracking cookies")},
		{Key: "content-blocking-cross-site-cookies", Value: ptr.Ptr("Cross-site cookies")},
		{Key: "content-blocking-private-windows", Value: ptr.Ptr("Tracking content in Private Windows")},
		{Key: "content-blocking-all-windows-tracking-content", Value: ptr.Ptr("Tracking content in all windows")},
		{Key: "content-blocking-all-third-party-cookies", Value: ptr.Ptr("All third-party cookies")},
		{Key: "content-blocking-all-cookies", Value: ptr.Ptr("All cookies")},
		{Key: "content-blocking-unvisited-cookies", Value: ptr.Ptr("Cookies from unvisited sites")},
		{Key: "content-blocking-cross-site-tracking-cookies", Value: ptr.Ptr("Cross-site tracking cookies")},
		{Key: "content-blocking-cross-site-tracking-cookies-plus-isolate", Value: ptr.Ptr("Cross-site tracking cookies, and isolate remaining cookies")},
		{Key: "content-blocking-cryptominers", Value: ptr.Ptr("Cryptominers")},
		{Key: "content-blocking-fingerprinters", Value: ptr.Ptr("Fingerprinters")},
		{Key: "content-blocking-reload-description", Value: ptr.Ptr("You will need to reload your tabs to apply these changes.")},
		{
			Key: "content-blocking-reload-tabs-button",
			Attrs: map[string]any{
				"label":     "Reload All Tabs",
				"accesskey": "R",
			},
		},
		{Key: "content-blocking-warning-title", Value: ptr.Ptr("Heads up!")},
		{
			Key:   "content-blocking-and-isolating-etp-warning-description",
			Value: ptr.Ptr("Blocking trackers and isolating cookies could impact the functionality of some sites. Reload a page with trackers to load all content."),
		},
		{Key: "content-blocking-warning-learn-how", Value: ptr.Ptr("Learn how")},
		{
			Key: "enhanced-tracking-protection-setting-custom",
			Attrs: map[string]any{
				"label":     "Custom",
				"accesskey": "C",
			},
		},
		{Key: "content-blocking-expand-section", Attrs: map[string]any{"tooltiptext": "More information"}},
		{Key: "content-blocking-etp-custom-desc", Value: ptr.Ptr("Choose which trackers and scripts to block.")},
		{
			Key: "content-blocking-cookies-label",
			Attrs: map[string]any{
				"label":     "Cookies",
				"accesskey": "C",
			},
		},
		{Key: "sitedata-option-block-cross-site-trackers", Attrs: map[string]any{"label": "Cross-site trackers"}},
		{
			Key: "sitedata-option-block-cross-site-and-social-media-trackers-plus-isolate",
			Attrs: map[string]any{
				"label": "Cross-site and social media trackers, and isolate remaining cookies",
			},
		},
		{Key: "sitedata-option-block-unvisited", Attrs: map[string]any{"label": "Cookies from unvisited websites"}},
		{Key: "sitedata-option-block-all-third-party", Attrs: map[string]any{"label": "All third-party cookies (may cause websites to break)"}},
		{Key: "sitedata-option-block-all", Attrs: map[string]any{"label": "All cookies (will cause websites to break)"}},
		{Key: "disable-extension", Attrs: map[string]any{"label": "Disable Extension"}},
		{
			Key: "content-blocking-tracking-content-label",
			Attrs: map[string]any{
				"label":     "Tracking content",
				"accesskey": "T",
			},
		},
		{
			Key: "content-blocking-option-private",
			Attrs: map[string]any{
				"label":     "Only in Private Windows",
				"accesskey": "p",
			},
		},
		{
			Key: "content-blocking-tracking-protection-option-all-windows",
			Attrs: map[string]any{
				"label":     "In all windows",
				"accesskey": "A",
			},
		},
		{Key: "content-blocking-tracking-protection-change-block-list", Value: ptr.Ptr("Change block list")},
		{
			Key: "content-blocking-cryptominers-label",
			Attrs: map[string]any{
				"label":     "Cryptominers",
				"accesskey": "y",
			},
		},
		{
			Key: "content-blocking-fingerprinters-label",
			Attrs: map[string]any{
				"label":     "Fingerprinters",
				"accesskey": "F",
			},
		},
		{Key: "content-blocking-reload-description", Value: ptr.Ptr("You will need to reload your tabs to apply these changes.")},
		{
			Key: "content-blocking-reload-tabs-button",
			Attrs: map[string]any{
				"label":     "Reload All Tabs",
				"accesskey": "R",
			},
		},
		{Key: "content-blocking-warning-title", Value: ptr.Ptr("Heads up!")},
		{
			Key:   "content-blocking-and-isolating-etp-warning-description",
			Value: ptr.Ptr("Blocking trackers and isolating cookies could impact the functionality of some sites. Reload a page with trackers to load all content."),
		},
		{Key: "content-blocking-warning-learn-how", Value: ptr.Ptr("Learn how")},
		{Key: "do-not-track-description", Value: ptr.Ptr("Send websites a “Do Not Track” signal that you don’t want to be tracked")},
		{Key: "do-not-track-learn-more", Value: ptr.Ptr("Learn more")},
		{Key: "do-not-track-option-always", Attrs: map[string]any{"label": "Always"}},
		{
			Key: "do-not-track-option-default-content-blocking-known",
			Attrs: map[string]any{
				"label": "Only when Nightly is set to block known trackers",
			},
		},
		{Key: "sitedata-header", Value: ptr.Ptr("Cookies and Site Data")},
		{Key: "sitedata-learn-more", Value: ptr.Ptr("Learn more")},
		{
			Key:   "sitedata-delete-on-close-private-browsing",
			Value: ptr.Ptr("In permanent private browsing mode, cookies and site data will always be cleared when Nightly is closed."),
		},
		{
			Key: "sitedata-delete-on-close",
			Attrs: map[string]any{
				"label":     "Delete cookies and site data when Nightly is closed",
				"accesskey": "c",
			},
		},
		{
			Key: "sitedata-clear",
			Attrs: map[string]any{
				"label":     "Clear Data…",
				"accesskey": "l",
			},
		},
		{
			Key: "sitedata-settings",
			Attrs: map[string]any{
				"label":     "Manage Data…",
				"accesskey": "M",
			},
		},
		{
			Key: "sitedata-cookies-exceptions",
			Attrs: map[string]any{
				"label":     "Manage Exceptions…",
				"accesskey": "x",
			},
		},
		{
			Key:   "pane-privacy-logins-and-passwords-header",
			Value: ptr.Ptr("Logins and Passwords"),
			Attrs: map[string]any{
				"searchkeywords": "Lockwise",
			},
		},
		{Key: "disable-extension", Attrs: map[string]any{"label": "Disable Extension"}},
		{
			Key: "forms-ask-to-save-logins",
			Attrs: map[string]any{
				"label":     "Ask to save logins and passwords for websites",
				"accesskey": "r",
			},
		},
		{
			Key: "forms-fill-logins-and-passwords",
			Attrs: map[string]any{
				"label":     "Autofill logins and passwords",
				"accesskey": "i",
			},
		},
		{
			Key: "forms-generate-passwords",
			Attrs: map[string]any{
				"label":     "Suggest and generate strong passwords",
				"accesskey": "u",
			},
		},
		{
			Key: "forms-exceptions",
			Attrs: map[string]any{
				"label":     "Exceptions…",
				"accesskey": "x",
			},
		},
		{
			Key: "forms-saved-logins",
			Attrs: map[string]any{
				"label":     "Saved Logins…",
				"accesskey": "L",
			},
		},
		{
			Key: "forms-breach-alerts",
			Attrs: map[string]any{
				"label":     "Show alerts about passwords for breached websites",
				"accesskey": "b",
			},
		},
		{Key: "forms-breach-alerts-learn-more-link", Value: ptr.Ptr("Learn more")},
		{
			Key: "forms-primary-pw-use",
			Attrs: map[string]any{
				"label":     "Use a Primary Password",
				"accesskey": "U",
			},
		},
		{Key: "forms-primary-pw-learn-more-link", Value: ptr.Ptr("Learn more")},
		{
			Key: "forms-primary-pw-change",
			Attrs: map[string]any{
				"label":     "Change Primary Password…",
				"accesskey": "P",
			},
		},
		{Key: "forms-primary-pw-former-name", Value: ptr.Ptr("Formerly known as Master Password")},
		{Key: "forms-primary-pw-fips-title", Value: ptr.Ptr("You are currently in FIPS mode. FIPS requires a non-empty Primary Password.")},
		{Key: "forms-master-pw-fips-desc", Value: ptr.Ptr("Password Change Failed")},
		{Key: "history-header", Value: ptr.Ptr("History")},
		{
			Key:   "history-remember-label",
			Value: ptr.Ptr("Nightly will"),
			Attrs: map[string]any{
				"accesskey": "w",
			},
		},
		{Key: "history-remember-option-all", Attrs: map[string]any{"label": "Remember history"}},
		{Key: "history-remember-option-never", Attrs: map[string]any{"label": "Never remember history"}},
		{Key: "history-remember-option-custom", Attrs: map[string]any{"label": "Use custom settings for history"}},
		{Key: "history-remember-description", Value: ptr.Ptr("Nightly will remember your browsing, download, form, and search history.")},
		{
			Key:   "history-dontremember-description",
			Value: ptr.Ptr("Nightly will use the same settings as private browsing, and will not remember any history as you browse the Web."),
		},
		{
			Key: "history-private-browsing-permanent",
			Attrs: map[string]any{
				"label":     "Always use private browsing mode",
				"accesskey": "p",
			},
		},
		{
			Key: "history-remember-browser-option",
			Attrs: map[string]any{
				"label":     "Remember browsing and download history",
				"accesskey": "b",
			},
		},
		{
			Key: "history-remember-search-option",
			Attrs: map[string]any{
				"label":     "Remember search and form history",
				"accesskey": "f",
			},
		},
		{
			Key: "history-clear-on-close-option",
			Attrs: map[string]any{
				"label":     "Clear history when Nightly closes",
				"accesskey": "r",
			},
		},
		{
			Key: "history-clear-button",
			Attrs: map[string]any{
				"label":     "Clear History…",
				"accesskey": "s",
			},
		},
		{
			Key: "history-clear-on-close-settings",
			Attrs: map[string]any{
				"label":     "Settings…",
				"accesskey": "t",
			},
		},
		{Key: "addressbar-header", Value: ptr.Ptr("Address Bar")},
		{Key: "addressbar-suggest", Value: ptr.Ptr("When using the address bar, suggest")},
		{
			Key: "addressbar-locbar-history-option",
			Attrs: map[string]any{
				"label":     "Browsing history",
				"accesskey": "h",
			},
		},
		{
			Key: "addressbar-locbar-bookmarks-option",
			Attrs: map[string]any{
				"label":     "Bookmarks",
				"accesskey": "k",
			},
		},
		{
			Key: "addressbar-locbar-openpage-option",
			Attrs: map[string]any{
				"label":     "Open tabs",
				"accesskey": "O",
			},
		},
		{
			Key: "addressbar-locbar-topsites-option",
			Attrs: map[string]any{
				"label":     "Top sites",
				"accesskey": "T",
			},
		},
		{
			Key: "addressbar-locbar-engines-option",
			Attrs: map[string]any{
				"label":     "Search engines",
				"accesskey": "a",
			},
		},
		{Key: "addressbar-suggestions-settings", Value: ptr.Ptr("Change preferences for search engine suggestions")},
		{Key: "permissions-header", Value: ptr.Ptr("Permissions")},
		{Key: "permissions-header", Value: ptr.Ptr("Permissions")},
		{Key: "permissions-location", Value: ptr.Ptr("Location")},
		{
			Key: "permissions-location-settings",
			Attrs: map[string]any{
				"label":     "Settings…",
				"accesskey": "t",
			},
		},
		{Key: "permissions-camera", Value: ptr.Ptr("Camera")},
		{
			Key: "permissions-camera-settings",
			Attrs: map[string]any{
				"label":     "Settings…",
				"accesskey": "t",
			},
		},
		{Key: "permissions-microphone", Value: ptr.Ptr("Microphone")},
		{
			Key: "permissions-microphone-settings",
			Attrs: map[string]any{
				"label":     "Settings…",
				"accesskey": "t",
			},
		},
		{Key: "permissions-notification", Value: ptr.Ptr("Notifications")},
		{Key: "permissions-notification-link", Value: ptr.Ptr("Learn more")},
		{
			Key: "permissions-notification-settings",
			Attrs: map[string]any{
				"label":     "Settings…",
				"accesskey": "t",
			},
		},
		{Key: "permissions-autoplay", Value: ptr.Ptr("Autoplay")},
		{
			Key: "permissions-autoplay-settings",
			Attrs: map[string]any{
				"label":     "Settings…",
				"accesskey": "t",
			},
		},
		{Key: "permissions-xr", Value: ptr.Ptr("Virtual Reality")},
		{
			Key: "permissions-xr-settings",
			Attrs: map[string]any{
				"label":     "Settings…",
				"accesskey": "t",
			},
		},
		{
			Key: "permissions-block-popups",
			Attrs: map[string]any{
				"label":     "Block pop-up windows",
				"accesskey": "B",
			},
		},
		{
			Key: "permissions-block-popups-exceptions",
			Attrs: map[string]any{
				"label":     "Exceptions…",
				"accesskey": "E",
			},
		},
		{
			Key: "permissions-addon-install-warning",
			Attrs: map[string]any{
				"label":     "Warn you when websites try to install add-ons",
				"accesskey": "W",
			},
		},
		{
			Key: "permissions-addon-exceptions",
			Attrs: map[string]any{
				"label":     "Exceptions…",
				"accesskey": "E",
			},
		},
		{Key: "collection-header", Value: ptr.Ptr("Nightly Data Collection and Use")},
		{Key: "collection-header", Value: ptr.Ptr("Nightly Data Collection and Use")},
		{
			Key:   "collection-description",
			Value: ptr.Ptr("We strive to provide you with choices and collect only what we need to provide and improve Nightly for everyone. We always ask permission before receiving personal information."),
		},
		{Key: "collection-privacy-notice", Value: ptr.Ptr("Privacy Notice")},
		{
			Key:   "collection-health-report-telemetry-disabled",
			Value: ptr.Ptr("You’re no longer allowing Mozilla to capture technical and interaction data. All past data will be deleted within 30 days."),
		},
		{Key: "collection-health-report-telemetry-disabled-link", Value: ptr.Ptr("Learn more")},
		{
			Key: "collection-health-report",
			Attrs: map[string]any{
				"label":     "Allow Nightly to send technical and interaction data to Mozilla",
				"accesskey": "r",
			},
		},
		{Key: "collection-health-report-link", Value: ptr.Ptr("Learn more")},
		{Key: "addon-recommendations", Attrs: map[string]any{"label": "Allow Nightly to make personalized extension recommendations"}},
		{Key: "addon-recommendations-link", Value: ptr.Ptr("Learn more")},
		{Key: "collection-health-report-disabled", Value: ptr.Ptr("Data reporting is disabled for this build configuration")},
		{Key: "collection-studies", Attrs: map[string]any{"label": "Allow Nightly to install and run studies"}},
		{Key: "collection-studies-link", Value: ptr.Ptr("View Nightly studies")},
		{
			Key: "collection-backlogged-crash-reports",
			Attrs: map[string]any{
				"label":     "Allow Nightly to send backlogged crash reports on your behalf",
				"accesskey": "c",
			},
		},
		{Key: "collection-backlogged-crash-reports-link", Value: ptr.Ptr("Learn more")},
		{Key: "security-header", Value: ptr.Ptr("Security")},
		{Key: "security-browsing-protection", Value: ptr.Ptr("Deceptive Content and Dangerous Software Protection")},
		{
			Key: "security-enable-safe-browsing",
			Attrs: map[string]any{
				"label":     "Block dangerous and deceptive content",
				"accesskey": "B",
			},
		},
		{Key: "security-enable-safe-browsing-link", Value: ptr.Ptr("Learn more")},
		{
			Key: "security-block-downloads",
			Attrs: map[string]any{
				"label":     "Block dangerous downloads",
				"accesskey": "d",
			},
		},
		{
			Key: "security-block-uncommon-software",
			Attrs: map[string]any{
				"label":     "Warn you about unwanted and uncommon software",
				"accesskey": "c",
			},
		},
		{Key: "certs-header", Value: ptr.Ptr("Certificates")},
		{
			Key: "certs-enable-ocsp",
			Attrs: map[string]any{
				"label":     "Query OCSP responder servers to confirm the current validity of certificates",
				"accesskey": "Q",
			},
		},
		{
			Key: "certs-view",
			Attrs: map[string]any{
				"label":     "View Certificates…",
				"accesskey": "C",
			},
		},
		{
			Key: "certs-devices",
			Attrs: map[string]any{
				"label":     "Security Devices…",
				"accesskey": "D",
			},
		},
		{Key: "httpsonly-header", Value: ptr.Ptr("HTTPS-Only Mode")},
		{
			Key:   "httpsonly-description",
			Value: ptr.Ptr("HTTPS provides a secure, encrypted connection between Nightly and the websites you visit. Most websites support HTTPS, and if HTTPS-Only Mode is enabled, then Nightly will upgrade all connections to HTTPS."),
		},
		{Key: "httpsonly-learn-more", Value: ptr.Ptr("Learn more")},
		{Key: "httpsonly-radio-enabled", Attrs: map[string]any{"label": "Enable HTTPS-Only Mode in all windows"}},
		{Key: "httpsonly-radio-enabled-pbm", Attrs: map[string]any{"label": "Enable HTTPS-Only Mode in private windows only"}},
		{Key: "httpsonly-radio-disabled", Attrs: map[string]any{"label": "Don’t enable HTTPS-Only Mode"}},
		{Key: "pane-experimental-title", Value: ptr.Ptr("Nightly Experiments")},
		{Key: "pane-experimental-subtitle", Value: ptr.Ptr("Proceed with Caution")},
		{
			Key: "pane-experimental-reset",
			Attrs: map[string]any{
				"label":     "Restore Defaults",
				"accesskey": "R",
			},
		},
		{Key: "pane-experimental-search-results-header", Value: ptr.Ptr("Nightly Experiments: Proceed with Caution")},
		{
			Key:   "pane-experimental-description",
			Value: ptr.Ptr("Changing advanced configuration preferences can impact Nightly performance or security."),
		},
		{Key: "pane-sync-title2", Value: ptr.Ptr("Sync")},
		{Key: "sync-signedout-caption", Value: ptr.Ptr("Take Your Web With You")},
		{
			Key:   "sync-signedout-description",
			Value: ptr.Ptr("Synchronize your bookmarks, history, tabs, passwords, add-ons, and preferences across all your devices."),
		},
		{
			Key: "sync-signedout-account-signin2",
			Attrs: map[string]any{
				"label":     "Sign in to Sync…",
				"accesskey": "i",
			},
		},
		{
			Key:   "sync-mobile-promo",
			Value: ptr.Ptr("Download Firefox for <img data-l10n-name=\"android-icon\"/> <a data-l10n-name=\"android-link\">Android</a> or <img data-l10n-name=\"ios-icon\"/> <a data-l10n-name=\"ios-link\">iOS</a> to sync with your mobile device."),
		},
		{Key: "pane-sync-title2", Value: ptr.Ptr("Sync")},
		{Key: "sync-profile-picture", Attrs: map[string]any{"tooltiptext": "Change profile picture"}},
		{
			Key: "sync-sign-out",
			Attrs: map[string]any{
				"label":     "Sign Out…",
				"accesskey": "g",
			},
		},
		{
			Key:   "sync-manage-account",
			Value: ptr.Ptr("Manage account"),
			Attrs: map[string]any{
				"accesskey": "o",
			},
		},
		{
			Key: "sync-signedin-unverified",
			Args: map[string]any{
				"email": "",
			},
			Value: ptr.Ptr(" is not verified."),
		},
		{
			Key: "sync-resend-verification",
			Attrs: map[string]any{
				"label":     "Resend Verification",
				"accesskey": "d",
			},
		},
		{
			Key: "sync-remove-account",
			Attrs: map[string]any{
				"label":     "Remove Account",
				"accesskey": "R",
			},
		},
		{
			Key: "sync-signedin-login-failure",
			Args: map[string]any{
				"email": "",
			},
			Value: ptr.Ptr("Please sign in to reconnect "),
		},
		{
			Key: "sync-sign-in",
			Attrs: map[string]any{
				"label":     "Sign in",
				"accesskey": "g",
			},
		},
		{
			Key: "sync-remove-account",
			Attrs: map[string]any{
				"label":     "Remove Account",
				"accesskey": "R",
			},
		},
		{Key: "sync-device-name-header", Value: ptr.Ptr("Device Name")},
		{
			Key: "sync-device-name-change",
			Attrs: map[string]any{
				"label":     "Change Device Name…",
				"accesskey": "h",
			},
		},
		{
			Key: "sync-device-name-cancel",
			Attrs: map[string]any{
				"label":     "Cancel",
				"accesskey": "n",
			},
		},
		{
			Key: "sync-device-name-save",
			Attrs: map[string]any{
				"label":     "Save",
				"accesskey": "v",
			},
		},
		{Key: "prefs-syncing-off", Value: ptr.Ptr("Syncing: OFF")},
		{
			Key:   "prefs-sync-offer-setup-label",
			Value: ptr.Ptr("Synchronize your bookmarks, history, tabs, passwords, add-ons, and preferences across all your devices."),
		},
		{
			Key: "prefs-sync-setup",
			Attrs: map[string]any{
				"label":     "Set Up Sync…",
				"accesskey": "S",
			},
		},
		{Key: "prefs-syncing-on", Value: ptr.Ptr("Syncing: ON")},
		{
			Key: "prefs-sync-now",
			Attrs: map[string]any{
				"labelnotsyncing":     "Sync Now",
				"accesskeynotsyncing": "N",
				"labelsyncing":        "Syncing…",
			},
		},
		{Key: "sync-currently-syncing-heading", Value: ptr.Ptr("You are currently syncing these items:")},
		{Key: "sync-currently-syncing-bookmarks", Value: ptr.Ptr("Bookmarks")},
		{Key: "sync-currently-syncing-history", Value: ptr.Ptr("History")},
		{Key: "sync-currently-syncing-tabs", Value: ptr.Ptr("Open tabs")},
		{Key: "sync-currently-syncing-logins-passwords", Value: ptr.Ptr("Logins and passwords")},
		{Key: "sync-currently-syncing-addresses", Value: ptr.Ptr("Addresses")},
		{Key: "sync-currently-syncing-creditcards", Value: ptr.Ptr("Credit cards")},
		{Key: "sync-currently-syncing-addons", Value: ptr.Ptr("Add-ons")},
		{Key: "sync-currently-syncing-prefs", Value: ptr.Ptr("Preferences")},
		{
			Key: "sync-change-options",
			Attrs: map[string]any{
				"label":     "Change…",
				"accesskey": "C",
			},
		},
		{Key: "sync-connect-another-device", Value: ptr.Ptr("Connect another device")},
		{Key: "experimental-features-abouthome-startup-cache", Attrs: map[string]any{"label": "about:home startup cache"}},
		{Key: "experimental-features-cookie-samesite-lax-by-default2", Attrs: map[string]any{"label": "Cookies: SameSite=Lax by default"}},
		{
			Key: "experimental-features-cookie-samesite-none-requires-secure2",
			Attrs: map[string]any{
				"label": "Cookies: SameSite=None requires secure attribute",
			},
		},
		{Key: "experimental-features-cookie-samesite-schemeful", Attrs: map[string]any{"label": "Cookies: Schemeful SameSite"}},
		{Key: "experimental-features-css-constructable-stylesheets", Attrs: map[string]any{"label": "CSS: Constructable Stylesheets"}},
		{Key: "experimental-features-css-focus-visible", Attrs: map[string]any{"label": "CSS: Pseudo-class: :focus-visible"}},
		{Key: "experimental-features-css-masonry2", Attrs: map[string]any{"label": "CSS: Masonry Layout"}},
		{Key: "experimental-features-devtools-color-scheme-simulation", Attrs: map[string]any{"label": "Developer Tools: Color Scheme Simulation"}},
		{Key: "experimental-features-devtools-compatibility-panel", Attrs: map[string]any{"label": "Developer Tools: Compatibility Panel"}},
		{
			Key: "experimental-features-devtools-execution-context-selector",
			Attrs: map[string]any{
				"label": "Developer Tools: Execution Context Selector",
			},
		},
		{
			Key: "experimental-features-devtools-serviceworker-debugger-support",
			Attrs: map[string]any{
				"label": "Developer Tools: Service Worker debugging",
			},
		},
		{Key: "experimental-features-fission", Attrs: map[string]any{"label": "Fission (Site Isolation)"}},
		{Key: "experimental-features-http3", Attrs: map[string]any{"label": "HTTP/3 protocol"}},
		{Key: "experimental-features-media-avif", Attrs: map[string]any{"label": "Media: AVIF"}},
		{Key: "experimental-features-multi-pip", Attrs: map[string]any{"label": "Multiple Picture-in-Picture Support"}},
		{Key: "experimental-features-print-preview-tab-modal", Attrs: map[string]any{"label": "Print Preview Redesign"}},
		{Key: "experimental-features-web-api-beforeinput", Attrs: map[string]any{"label": "Web API: beforeinput Event"}},
		{Key: "experimental-features-web-api-inputmode", Attrs: map[string]any{"label": "Web API: inputmode"}},
		{Key: "experimental-features-web-api-link-preload", Attrs: map[string]any{"label": "Web API: <link rel=\"preload\">"}},
		{Key: "experimental-features-web-gpu2", Attrs: map[string]any{"label": "Web API: WebGPU"}},
		{Key: "experimental-features-webrtc-global-mute-toggles", Attrs: map[string]any{"label": "WebRTC Global Mute Toggles"}},
		{
			Key: "use-current-pages",
			Args: map[string]any{
				"tabCount": "1",
			},
			Attrs: map[string]any{
				"label":     "Use Current Page",
				"accesskey": "C",
			},
		},
		{Key: "home-prefs-content-header", Value: ptr.Ptr("Firefox Home Content")},
		{Key: "home-prefs-content-description", Value: ptr.Ptr("Choose what content you want on your Firefox Home screen.")},
		{Key: "home-prefs-search-header", Attrs: map[string]any{"label": "Web Search"}},
		{Key: "home-prefs-topsites-header", Attrs: map[string]any{"label": "Top Sites"}},
		{Key: "home-prefs-topsites-description", Value: ptr.Ptr("The sites you visit most")},
		{
			Key: "home-prefs-sections-rows-option",
			Args: map[string]any{
				"num": "1",
			},
			Attrs: map[string]any{
				"label": "1 row",
			},
		},
		{
			Key: "home-prefs-sections-rows-option",
			Args: map[string]any{
				"num": "2",
			},
			Attrs: map[string]any{
				"label": "2 rows",
			},
		},
		{
			Key: "home-prefs-sections-rows-option",
			Args: map[string]any{
				"num": "3",
			},
			Attrs: map[string]any{
				"label": "3 rows",
			},
		},
		{
			Key: "home-prefs-sections-rows-option",
			Args: map[string]any{
				"num": "4",
			},
			Attrs: map[string]any{
				"label": "4 rows",
			},
		},
		{
			Key: "home-prefs-recommended-by-header",
			Args: map[string]any{
				"provider": "Pocket",
			},
			Attrs: map[string]any{
				"label": "Recommended by Pocket",
			},
		},
		{Key: "home-prefs-recommended-by-learn-more", Value: ptr.Ptr("How it works")},
		{
			Key: "home-prefs-recommended-by-description-update",
			Args: map[string]any{
				"provider": "Pocket",
			},
			Value: ptr.Ptr("Exceptional content from across the web, curated by Pocket"),
		},
		{Key: "home-prefs-recommended-by-option-sponsored-stories", Attrs: map[string]any{"label": "Sponsored Stories"}},
		{Key: "home-prefs-highlights-header", Attrs: map[string]any{"label": "Highlights"}},
		{Key: "home-prefs-highlights-description", Value: ptr.Ptr("A selection of sites that you’ve saved or visited")},
		{
			Key: "home-prefs-sections-rows-option",
			Args: map[string]any{
				"num": "1",
			},
			Attrs: map[string]any{
				"label": "1 row",
			},
		},
		{
			Key: "home-prefs-sections-rows-option",
			Args: map[string]any{
				"num": "2",
			},
			Attrs: map[string]any{
				"label": "2 rows",
			},
		},
		{
			Key: "home-prefs-sections-rows-option",
			Args: map[string]any{
				"num": "3",
			},
			Attrs: map[string]any{
				"label": "3 rows",
			},
		},
		{
			Key: "home-prefs-sections-rows-option",
			Args: map[string]any{
				"num": "4",
			},
			Attrs: map[string]any{
				"label": "4 rows",
			},
		},
		{Key: "home-prefs-highlights-option-visited-pages", Attrs: map[string]any{"label": "Visited Pages"}},
		{Key: "home-prefs-highlights-options-bookmarks", Attrs: map[string]any{"label": "Bookmarks"}},
		{Key: "home-prefs-highlights-option-most-recent-download", Attrs: map[string]any{"label": "Most Recent Download"}},
		{Key: "home-prefs-highlights-option-saved-to-pocket", Attrs: map[string]any{"label": "Pages Saved to Pocket"}},
		{Key: "home-prefs-snippets-header", Attrs: map[string]any{"label": "Snippets"}},
		{Key: "home-prefs-snippets-description", Value: ptr.Ptr("Updates from Mozilla and Firefox")},
		{Key: "sitedata-option-block-cross-site-and-social-media-trackers", Attrs: map[string]any{"label": "Cross-site and social media trackers"}},
		{Key: "sitedata-total-size-calculating", Value: ptr.Ptr("Calculating site data and cache size…")},
		{Key: "containers-preferences-button", Attrs: map[string]any{"label": "Preferences"}},
		{Key: "containers-remove-button", Attrs: map[string]any{"label": "Remove"}},
		{Key: "containers-preferences-button", Attrs: map[string]any{"label": "Preferences"}},
		{Key: "containers-remove-button", Attrs: map[string]any{"label": "Remove"}},
		{Key: "containers-preferences-button", Attrs: map[string]any{"label": "Preferences"}},
		{Key: "containers-remove-button", Attrs: map[string]any{"label": "Remove"}},
		{Key: "containers-preferences-button", Attrs: map[string]any{"label": "Preferences"}},
		{Key: "containers-remove-button", Attrs: map[string]any{"label": "Remove"}},
		{Key: "experimental-features-abouthome-startup-cache", Attrs: map[string]any{"label": "about:home startup cache"}},
		{
			Key:   "experimental-features-abouthome-startup-cache-description",
			Value: ptr.Ptr("A cache for the initial about:home document that is loaded by default at startup. The purpose of the cache is to improve startup performance."),
		},
		{Key: "experimental-features-cookie-samesite-lax-by-default2", Attrs: map[string]any{"label": "Cookies: SameSite=Lax by default"}},
		{
			Key:   "experimental-features-cookie-samesite-lax-by-default2-description",
			Value: ptr.Ptr("Treat cookies as “SameSite=Lax” by default if no “SameSite” attribute is specified. Developers must opt-in to the current status quo of unrestricted use by explicitly asserting “SameSite=None”."),
		},
		{
			Key: "experimental-features-cookie-samesite-none-requires-secure2",
			Attrs: map[string]any{
				"label": "Cookies: SameSite=None requires secure attribute",
			},
		},
		{
			Key:   "experimental-features-cookie-samesite-none-requires-secure2-description",
			Value: ptr.Ptr("Cookies with “SameSite=None” attribute require the secure attribute. This feature requires “Cookies: SameSite=Lax by default”."),
		},
		{Key: "experimental-features-cookie-samesite-schemeful", Attrs: map[string]any{"label": "Cookies: Schemeful SameSite"}},
		{
			Key:   "experimental-features-cookie-samesite-schemeful-description",
			Value: ptr.Ptr("Treat cookies from the same domain, but with different schemes (e.g. http://example.com and https://example.com) as cross-site instead of same-site. Improves security, but potentially introduces breakage."),
		},
		{Key: "experimental-features-css-constructable-stylesheets", Attrs: map[string]any{"label": "CSS: Constructable Stylesheets"}},
		{
			Key:   "experimental-features-css-constructable-stylesheets-description",
			Value: ptr.Ptr("The addition of a constructor to the <a data-l10n-name=\"mdn-cssstylesheet\">CSSStyleSheet</a> interface as well as a variety of related changes makes it possible to directly create new stylesheets without having to add the sheet to the HTML. This makes it much easier to create reusable stylesheets for use with <a data-l10n-name=\"mdn-shadowdom\">Shadow DOM</a>. See <a data-l10n-name=\"bugzilla\">bug 1520690</a> for more details."),
		},
		{Key: "experimental-features-css-masonry2", Attrs: map[string]any{"label": "CSS: Masonry Layout"}},
		{
			Key:   "experimental-features-css-masonry-description",
			Value: ptr.Ptr("Enables support for the experimental CSS Masonry Layout feature. See the <a data-l10n-name=\"explainer\">explainer</a> for a high level description of the feature. To provide feedback, please comment in <a data-l10n-name=\"w3c-issue\">this GitHub issue</a> or <a data-l10n-name=\"bug\">this bug</a>."),
		},
		{Key: "experimental-features-css-focus-visible", Attrs: map[string]any{"label": "CSS: Pseudo-class: :focus-visible"}},
		{
			Key:   "experimental-features-css-focus-visible-description",
			Value: ptr.Ptr("Allows focus styles to be applied to elements like buttons and form controls, only when they are focused using the keyboard (e.g. when tabbing between elements), and not when they are focused using a mouse or other pointing device. See <a data-l10n-name=\"bugzilla\">bug 1617600</a> for more details."),
		},
		{Key: "experimental-features-devtools-color-scheme-simulation", Attrs: map[string]any{"label": "Developer Tools: Color Scheme Simulation"}},
		{
			Key:   "experimental-features-devtools-color-scheme-simulation-description",
			Value: ptr.Ptr("Adds an option to simulate different color schemes allowing you to test <a data-l10n-name=\"mdn-preferscolorscheme\">@prefers-color-scheme</a> media queries. Using this media query lets your stylesheet respond to whether the user prefers a light or dark user interface. This feature lets you test your code without having to change settings in your browser (or operating system, if the browser follows a system-wide color scheme setting). See <a data-l10n-name=\"bugzilla1\">bug 1550804</a> and <a data-l10n-name=\"bugzilla2\">bug 1137699</a> for more details."),
		},
		{Key: "experimental-features-devtools-compatibility-panel", Attrs: map[string]any{"label": "Developer Tools: Compatibility Panel"}},
		{
			Key:   "experimental-features-devtools-compatibility-panel-description",
			Value: ptr.Ptr("A side panel for the Page Inspector that shows you information detailing your app’s cross-browser compatibility status. See <a data-l10n-name=\"bugzilla\">bug 1584464</a> for more details."),
		},
		{
			Key: "experimental-features-devtools-execution-context-selector",
			Attrs: map[string]any{
				"label": "Developer Tools: Execution Context Selector",
			},
		},
		{
			Key:   "experimental-features-devtools-execution-context-selector-description",
			Value: ptr.Ptr("This feature displays a button on the console’s command line that lets you change the context in which the expression you enter will be executed. See <a data-l10n-name=\"bugzilla1\">bug 1605154</a> and <a data-l10n-name=\"bugzilla2\">bug 1605153</a> for more details."),
		},
		{
			Key: "experimental-features-devtools-serviceworker-debugger-support",
			Attrs: map[string]any{
				"label": "Developer Tools: Service Worker debugging",
			},
		},
		{
			Key:   "experimental-features-devtools-serviceworker-debugger-support-description",
			Value: ptr.Ptr("Enables experimental support for Service Workers in the Debugger panel. This feature may slow the Developer Tools down and increase memory consumption."),
		},
		{Key: "experimental-features-fission", Attrs: map[string]any{"label": "Fission (Site Isolation)"}},
		{
			Key:   "experimental-features-fission-description",
			Value: ptr.Ptr("Fission (site isolation) is an experimental feature in Nightly to provide an additional layer of defense against security bugs. By isolating each site into a separate process, Fission makes it harder for malicious websites to get access to information from other pages you are visiting. This is a major architectural change in Nightly and we appreciate you testing and reporting any issues you might encounter. For more details, see <a data-l10n-name=\"wiki\">the wiki</a>."),
		},
		{Key: "experimental-features-http3", Attrs: map[string]any{"label": "HTTP/3 protocol"}},
		{Key: "experimental-features-http3-description", Value: ptr.Ptr("Experimental support for the HTTP/3 protocol.")},
		{Key: "experimental-features-media-avif", Attrs: map[string]any{"label": "Media: AVIF"}},
		{
			Key:   "experimental-features-media-avif-description",
			Value: ptr.Ptr("With this feature enabled, Nightly supports the AV1 Image File (AVIF) format. This is a still image file format that leverages the capabilities of the AV1 video compression algorithms to reduce image size. See <a data-l10n-name=\"bugzilla\">bug 1443863</a> for more details."),
		},
		{Key: "experimental-features-multi-pip", Attrs: map[string]any{"label": "Multiple Picture-in-Picture Support"}},
		{
			Key:   "experimental-features-multi-pip-description",
			Value: ptr.Ptr("Experimental support for allowing multiple Picture-in-Picture windows to be open at the same time."),
		},
		{Key: "experimental-features-print-preview-tab-modal", Attrs: map[string]any{"label": "Print Preview Redesign"}},
		{
			Key:   "experimental-features-print-preview-tab-modal-description",
			Value: ptr.Ptr("Introduces the redesigned print preview and makes print preview available on macOS. This potentially introduces breakage and does not include all print-related settings. To access all print-related settings, select “Print using the system dialog…” from within the Print panel."),
		},
		{Key: "experimental-features-web-api-link-preload", Attrs: map[string]any{"label": "Web API: <link rel=\"preload\">"}},
		{
			Key:   "experimental-features-web-api-link-preload-description",
			Value: ptr.Ptr("The <a data-l10n-name=\"rel\">rel</a> attribute with value <code>\"preload\"</code> on a <a data-l10n-name=\"link\">&lt;link&gt;</a> element is intended to help provide performance gains by letting you download resources earlier in the page lifecycle, ensuring that they’re available earlier and are less likely to block page rendering. Read <a data-l10n-name=\"readmore\">“Preloading content with <code>rel=\"preload\"</code>”</a> or see <a data-l10n-name=\"bugzilla\">bug 1583604</a> for more details."),
		},
		{Key: "experimental-features-web-api-beforeinput", Attrs: map[string]any{"label": "Web API: beforeinput Event"}},
		{
			Key:   "experimental-features-web-api-beforeinput-description",
			Value: ptr.Ptr("The global <a data-l10n-name=\"mdn-beforeinput\">beforeinput</a> event is fired on an <a data-l10n-name=\"mdn-input\">&lt;input&gt;</a> and <a data-l10n-name=\"mdn-textarea\">&lt;textarea&gt;</a> elements, or any element whose <a data-l10n-name=\"mdn-contenteditable\">contenteditable</a> attribute is enabled, immediately before the element’s value changes. The event allows web apps to override the browser’s default behavior for user interaction, e.g., web apps can cancel user input only for specific characters or can modify pasting styled text only with approved styles."),
		},
		{Key: "experimental-features-web-api-inputmode", Attrs: map[string]any{"label": "Web API: inputmode"}},
		{
			Key:   "experimental-features-web-api-inputmode-description",
			Value: ptr.Ptr("Our implementation of the <a data-l10n-name=\"mdn-inputmode\">inputmode</a> global attribute has been updated as per <a data-l10n-name=\"whatwg\">the WHATWG specification</a>, but we still need to make other changes too, like making it available on contenteditable content. See <a data-l10n-name=\"bugzilla\">bug 1205133</a> for more details."),
		},
		{Key: "experimental-features-web-gpu2", Attrs: map[string]any{"label": "Web API: WebGPU"}},
		{
			Key:   "experimental-features-web-gpu-description2",
			Value: ptr.Ptr("This new API provides low-level support for performing computation and graphics rendering using the <a data-l10n-name=\"wikipedia\">Graphics Processing Unit (GPU)</a> of the user’s device or computer. The <a data-l10n-name=\"spec\">specification</a> is still a work-in-progress. See <a data-l10n-name=\"bugzilla\">bug 1602129</a> for more details."),
		},
		{Key: "experimental-features-webrtc-global-mute-toggles", Attrs: map[string]any{"label": "WebRTC Global Mute Toggles"}},
		{
			Key:   "experimental-features-webrtc-global-mute-toggles-description",
			Value: ptr.Ptr("Add controls to the WebRTC global sharing indicator that allow users to globally mute their microphone and camera feeds."),
		},
		{
			Key: "sitedata-total-size",
			Args: map[string]any{
				"value": "2.0",
				"unit":  "MB",
			},
			Value: ptr.Ptr("Your stored cookies, site data, and cache are currently using 2.0 MB of disk space."),
		},
	},
}

var EmptyResourceOneLocaleScenario = Scenario{
	Name:   "empty_resource_one_locale",
	Locale: "en-US",
	FileSources: []FileSource{
		{Name: "browser", PathScheme: "browser/{locale}/"},
		{Name: "empty", PathScheme: "empty-resource/{locale}/"},
	},
	Queries: []Query{
		{Key: "history-section-label", Value: ptr.Ptr("History")},
		{Key: "empty-one", Value: ptr.Ptr("pusty")},
	},
}

var EmptyResourceAllLocalesScenario = Scenario{
	Name:   "empty_resource_all_locales",
	Locale: "en-US",
	FileSources: []FileSource{
		{Name: "browser", PathScheme: "browser/{locale}/"},
		{Name: "empty", PathScheme: "empty-resource/{locale}/"},
	},
	Queries: []Query{
		{Key: "history-section-label", Value: ptr.Ptr("History")},
		{Key: "empty-all", Value: ptr.Ptr("empty-all")},
	},
}

var MissingOptionalOneLocaleScenario = Scenario{
	Name:   "missing_optional_one_locale",
	Locale: "en-US",
	FileSources: []FileSource{
		{Name: "browser", PathScheme: "browser/{locale}/"},
		{Name: "missing", PathScheme: "missing-resource/{locale}/"},
	},
	Queries: []Query{
		{Key: "history-section-label", Value: ptr.Ptr("History")},
		{Key: "missing-one", Value: ptr.Ptr("zaginiony")},
	},
}

var MissingOptionalAllLocalesScenario = Scenario{
	Name:   "missing_optional_all_locales",
	Locale: "en-US",
	FileSources: []FileSource{
		{Name: "browser", PathScheme: "browser/{locale}/"},
		{Name: "missing", PathScheme: "missing-resource/{locale}/"},
	},
	Queries: []Query{
		{Key: "history-section-label", Value: ptr.Ptr("History")},
		{Key: "missing-one", Value: ptr.Ptr("zaginiony")},
		{Key: "missing-all", Value: ptr.Ptr("missing-all")},
	},
}

var MissingRequiredOneLocaleScenario = Scenario{
	Name:   "missing_required_one_locale",
	Locale: "en-US",
	FileSources: []FileSource{
		{Name: "browser", PathScheme: "browser/{locale}/"},
		{Name: "missing", PathScheme: "missing-resource/{locale}/"},
	},
	Queries: []Query{
		{Key: "history-section-label", Value: ptr.Ptr("Historia")},
		{Key: "missing-one", Value: ptr.Ptr("zaginiony")},
	},
}

var MissingRequiredAllLocalesScenario = Scenario{
	Name:   "missing_required_all_locales",
	Locale: "en-US",
	FileSources: []FileSource{
		{Name: "browser", PathScheme: "browser/{locale}/"},
		{Name: "missing", PathScheme: "missing-resource/{locale}/"},
	},
	Queries: []Query{
		{Key: "history-section-label", Value: ptr.Ptr("history-section-label")},
		{Key: "missing-one", Value: ptr.Ptr("missing-one")},
		{Key: "missing-all", Value: ptr.Ptr("missing-all")},
	},
}
