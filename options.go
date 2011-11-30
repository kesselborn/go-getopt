package getopt

type Options []Option;

func (options Options) FindOption(optionString string) (option Option, found bool) {
  for _, cur := range options {
    if cur.ShortOpt() == optionString || cur.LongOpt() == optionString {
      option = cur
      found = true
      break
    }
  }

  return option, found
}
