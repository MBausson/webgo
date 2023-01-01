package Mime

func GetMime(ext string) string {
	m, found := filetype_to_mime[ext]

	if found {
		return m
	}

	return GetDefault()
}

func GetDefault() string {
	return filetype_to_mime["default"]
}

var filetype_to_mime map[string]string = map[string]string{
	"html":    "text/html",
	"css":     "text/css",
	"js":      "text/javascript",
	"json":    "application/json",
	"png":     "image/png",
	"jpg":     "image/jpeg",
	"jpeg":    "image/jpeg",
	"gif":     "image/gif",
	"svg":     "image/svg+xml",
	"ico":     "image/x-icon",
	"ttf":     "font/ttf",
	"woff":    "font/woff",
	"woff2":   "font/woff2",
	"eot":     "application/vnd.ms-fontobject",
	"otf":     "font/otf",
	"txt":     "text/plain",
	"pdf":     "application/pdf",
	"zip":     "application/zip",
	"rar":     "application/x-rar-compressed",
	"7z":      "application/x-7z-compressed",
	"gz":      "application/gzip",
	"tar":     "application/x-tar",
	"xml":     "application/xml",
	"mp3":     "audio/mpeg",
	"wav":     "audio/wav",
	"ogg":     "audio/ogg",
	"mp4":     "video/mp4",
	"webm":    "video/webm",
	"mkv":     "video/x-matroska",
	"avi":     "video/x-msvideo",
	"flv":     "video/x-flv",
	"mov":     "video/quicktime",
	"wmv":     "video/x-ms-wmv",
	"swf":     "application/x-shockwave-flash",
	"exe":     "application/x-msdownload",
	"msi":     "application/x-msdownload",
	"cab":     "application/vnd.ms-cab-compressed",
	"deb":     "application/x-debian-package",
	"rpm":     "application/x-redhat-package-manager",
	"apk":     "application/vnd.android.package-archive",
	"jar":     "application/java-archive",
	"war":     "application/java-archive",
	"ear":     "application/java-archive",
	"doc":     "application/msword",
	"docx":    "application/vnd.openxmlformats-officedocument.wordprocessingml.document",
	"xls":     "application/vnd.ms-excel",
	"xlsx":    "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet",
	"ppt":     "application/vnd.ms-powerpoint",
	"pptx":    "application/vnd.openxmlformats-officedocument.presentationml.presentation",
	"odt":     "application/vnd.oasis.opendocument.text",
	"ods":     "application/vnd.oasis.opendocument.spreadsheet",
	"odp":     "application/vnd.oasis.opendocument.presentation",
	"odg":     "application/vnd.oasis.opendocument.graphics",
	"odc":     "application/vnd.oasis.opendocument.chart",
	"odf":     "application/vnd.oasis.opendocument.formula",
	"odb":     "application/vnd.oasis.opendocument.database",
	"rtf":     "application/rtf",
	"csv":     "text/csv",
	"tsv":     "text/tab-separated-values",
	"ics":     "text/calendar",
	"rtx":     "text/richtext",
	"vcs":     "text/x-vcalendar",
	"vcf":     "text/x-vcard",
	"3gp":     "video/3gpp",
	"3g2":     "video/3gpp2",
	"abw":     "application/x-abiword",
	"arc":     "application/x-freearc",
	"azw":     "application/vnd.amazon.ebook",
	"bin":     "application/octet-stream",
	"bmp":     "image/bmp",
	"bz":      "application/x-bzip",
	"bz2":     "application/x-bzip2",
	"csh":     "application/x-csh",
	"conf":    "text/plain",
	"config":  "text/plain",
	"default": "text/plain",
}
