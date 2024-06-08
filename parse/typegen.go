// Code generated by "core generate -add-types"; DO NOT EDIT.

package parse

import (
	"cogentcore.org/core/types"
)

var _ = types.AddType(&types.Type{Name: "cogentcore.org/core/parse.FileState", IDName: "file-state", Doc: "FileState contains the full lexing and parsing state information for a given file.\nIt is the master state record for everything that happens in parse.  One of these\nshould be maintained for each file; texteditor.Buf has one as ParseState field.\n\nSeparate State structs are maintained for each stage (Lexing, PassTwo, Parsing) and\nthe final output of Parsing goes into the Ast and Syms fields.\n\nThe Src lexer.File field maintains all the info about the source file, and the basic\ntokenized version of the source produced initially by lexing and updated by the\nremaining passes.  It has everything that is maintained at a line-by-line level.", Fields: []types.Field{{Name: "Src", Doc: "the source to be parsed -- also holds the full lexed tokens"}, {Name: "LexState", Doc: "state for lexing"}, {Name: "TwoState", Doc: "state for second pass nesting depth and EOS matching"}, {Name: "ParseState", Doc: "state for parsing"}, {Name: "Ast", Doc: "ast output tree from parsing"}, {Name: "Syms", Doc: "symbols contained within this file -- initialized at start of parsing and created by AddSymbol or PushNewScope actions.  These are then processed after parsing by the language-specific code, via Lang interface."}, {Name: "ExtSyms", Doc: "External symbols that are entirely maintained in a language-specific way by the Lang interface code.  These are only here as a convenience and are not accessed in any way by the language-general parse code."}, {Name: "SymsMu", Doc: "mutex protecting updates / reading of Syms symbols"}, {Name: "WaitGp", Doc: "waitgroup for coordinating processing of other items"}, {Name: "AnonCtr", Doc: "anonymous counter -- counts up"}, {Name: "PathMap", Doc: "path mapping cache -- for other files referred to by this file, this stores the full path associated with a logical path (e.g., in go, the logical import path -> local path with actual files) -- protected for access from any thread"}}})

var _ = types.AddType(&types.Type{Name: "cogentcore.org/core/parse.FileStates", IDName: "file-states", Doc: "FileStates contains two FileState's: one is being processed while the\nother is being used externally.  The FileStates maintains\na common set of file information set in each of the FileState items when\nthey are used.", Fields: []types.Field{{Name: "Filename", Doc: "the filename"}, {Name: "Sup", Doc: "the known file type, if known (typically only known files are processed)"}, {Name: "BasePath", Doc: "base path for reporting file names -- this must be set externally e.g., by gide for the project root path"}, {Name: "DoneIndex", Doc: "index of the state that is done"}, {Name: "FsA", Doc: "one filestate"}, {Name: "FsB", Doc: "one filestate"}, {Name: "SwitchMu", Doc: "mutex locking the switching of Done vs. Proc states"}, {Name: "ProcMu", Doc: "mutex locking the parsing of Proc state -- reading states can happen fine with this locked, but no switching"}, {Name: "Meta", Doc: "extra meta data associated with this FileStates"}}})

var _ = types.AddType(&types.Type{Name: "cogentcore.org/core/parse.Lang", IDName: "lang", Doc: "Lang provides a general interface for language-specific management\nof the lexing, parsing, and symbol lookup process.\nThe parse lexer and parser machinery is entirely language-general\nbut specific languages may need specific ways of managing these\nprocesses, and processing their outputs, to best support the\nfeatures of those languages.  That is what this interface provides.\n\nEach language defines a type supporting this interface, which is\nin turn registered with the StdLangProperties map.  Each supported\nlanguage has its own .go file in this parse package that defines its\nown implementation of the interface and any other associated\nfunctionality.\n\nThe Lang is responsible for accessing the appropriate [Parser] for this\nlanguage (initialized and managed via LangSupport.OpenStandard() etc)\nand the [FileState] structure contains all the input and output\nstate information for a given file.\n\nThis interface is likely to evolve as we expand the range of supported\nlanguages."})

var _ = types.AddType(&types.Type{Name: "cogentcore.org/core/parse.LangDirOpts", IDName: "lang-dir-opts", Doc: "LangDirOpts provides options for Lang ParseDir method", Fields: []types.Field{{Name: "Subdirs", Doc: "process subdirectories -- otherwise not"}, {Name: "Rebuild", Doc: "rebuild the symbols by reprocessing from scratch instead of using cache"}, {Name: "Nocache", Doc: "do not update the cache with results from processing"}}})

var _ = types.AddType(&types.Type{Name: "cogentcore.org/core/parse.LangFlags", IDName: "lang-flags", Doc: "LangFlags are special properties of a given language"})

var _ = types.AddType(&types.Type{Name: "cogentcore.org/core/parse.LangProperties", IDName: "lang-properties", Doc: "LangProperties contains properties of languages supported by the Pi parser\nframework", Fields: []types.Field{{Name: "Known", Doc: "known language -- must be a supported one from Known list"}, {Name: "CommentLn", Doc: "character(s) that start a single-line comment -- if empty then multi-line comment syntax will be used"}, {Name: "CommentSt", Doc: "character(s) that start a multi-line comment or one that requires both start and end"}, {Name: "CommentEd", Doc: "character(s) that end a multi-line comment or one that requires both start and end"}, {Name: "Flags", Doc: "special properties for this language -- as an explicit list of options to make them easier to see and set in defaults"}, {Name: "Lang", Doc: "Lang interface for this language"}, {Name: "Parser", Doc: "parser for this language -- initialized in OpenStandard"}}})

var _ = types.AddType(&types.Type{Name: "cogentcore.org/core/parse.LangSupporter", IDName: "lang-supporter", Doc: "LangSupporter provides general support for supported languages.\ne.g., looking up lexers and parsers by name.\nAlso implements the lexer.LangLexer interface to provide access to other\nGuest Lexers"})

var _ = types.AddType(&types.Type{Name: "cogentcore.org/core/parse.Parser", IDName: "parser", Doc: "Parser is the overall parser for managing the parsing", Fields: []types.Field{{Name: "Lexer", Doc: "lexer rules for first pass of lexing file"}, {Name: "PassTwo", Doc: "second pass after lexing -- computes nesting depth and EOS finding"}, {Name: "Parser", Doc: "parser rules for parsing lexed tokens"}, {Name: "Filename", Doc: "file name for overall parser (not file being parsed!)"}, {Name: "ReportErrs", Doc: "if true, reports errors after parsing, to stdout"}, {Name: "ModTime", Doc: "when loaded from file, this is the modification time of the parser -- re-processes cache if parser is newer than cached files"}}})
