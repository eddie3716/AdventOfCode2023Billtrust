using System;
using System.Data.SqlClient;
using System.IO;

class Program
{
    static void Main(string[] args)
    {
        string server = args[0];
        string database = args[1];
        string fileId = args[2];

        string connectionString = $"Server={server};Database={database};Trusted_Connection=True;";

        string query = $"SELECT Content FROM [AR].[File] WHERE FileId = @FileId";

        using (SqlConnection connection = new SqlConnection(connectionString))
        {
            SqlCommand command = new SqlCommand(query, connection);
            command.Parameters.AddWithValue("@FileId", fileId);

            connection.Open();

            SqlDataReader reader = command.ExecuteReader();

            if (reader.Read())
            {
                byte[] data = (byte[])reader[0]; // Assuming the byte array is in the first column
                File.WriteAllBytes("output.txt", data);
            }
        }
    }
}