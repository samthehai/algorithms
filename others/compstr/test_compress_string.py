
from others.compstr.compress_string import CompressString
from nose.tools import assert_equal


class TestCompressString(object):
    def test_compress_string(self):
        cmpr_str = CompressString()

        assert_equal(cmpr_str.compress(""), "",)
        assert_equal(cmpr_str.compress("AABBCC"), "AABBCC",)
        assert_equal(cmpr_str.compress("AAABCCDDDDE"), "A3BC2D4E",)
        assert_equal(cmpr_str.compress("BAAACCDDDD"), "BA3C2D4", )
        assert_equal(cmpr_str.compress("AAABAACCDDDD"), "A3BA2C2D4",)

        print('Success: test_compress_string')


def main():
    test = TestCompressString()
    test.test_compress_string()


if __name__ == '__main__':
    main()
