using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace ConsoleApp1
{
    public static class StringExtensions
    {
        public static string BetterReverse(this string s)
        {
            return new string(s.Reverse().ToArray());
        }
    }

    
}
