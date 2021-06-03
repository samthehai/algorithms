from others import perfumations
from others.perfumations.is_permutation import Permutations
from nose.tools import assert_equal

class TestPermutation(object):
    def test_permutation(self):
        perm = Permutations()

        assert_equal(perm.is_permutation(None, 'foo'), False)
        assert_equal(perm.is_permutation('', 'foo'), False)
        assert_equal(perm.is_permutation('Nib', 'bin'), False)
        assert_equal(perm.is_permutation('act', 'cat'), True)
        assert_equal(perm.is_permutation('a ct', 'ca t'), True)

        print('Success: test_permutation')

def main():
    test = TestPermutation()
    test.test_permutation()

if __name__ == '__main__':
    main()
