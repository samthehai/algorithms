class CompressString(object):
    def compress(self, str):
        if str is None or not str:
            return str

        compstr, prev, count = '', str[0], 1

        for c in str:
            if c == prev:
                count += 1
            else:
                compstr += self._calc_partial_compstr(prev, count)
                prev = c
                count = 1

        compstr += self._calc_partial_compstr(prev, count)
        return compstr if len(compstr) < len(str) else str

    def _calc_partial_compstr(self, prev, count):
        return prev + (str(count) if count > 1 else '')
