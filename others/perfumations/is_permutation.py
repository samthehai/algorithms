from typing import DefaultDict


class Permutations(object):
    def is_permutation(self, str1, str2) -> bool:
        if str1 is None or str2 is None:
            return False

        if len(str1) != len(str2):
            return False

        m1 = DefaultDict(int)
        for c in str1:
            m1[c] += 1

        m2 = DefaultDict(int)
        for c in str2:
            m2[c] += 1

        return m1 == m2
